package auth

import (
	"backend/internal/utils"
	"backend/orm"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthHandler interface {
	FullfillLogin(c *gin.Context, dbConn *pgxpool.Conn, user goth.User)
	LogInUser(c *gin.Context)
	LogInUserCallback(c *gin.Context)
	LogOutUser(c *gin.Context)
}

type DefaultAuthHandler struct {
	AuthHandler
	dbPool *pgxpool.Pool
}

func NewDefaultAuthHandler(dbPool *pgxpool.Pool) *DefaultAuthHandler {
	return &DefaultAuthHandler{
		dbPool: dbPool,
	}
}

func (ah *DefaultAuthHandler) getConn(c *gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := ah.dbPool.Acquire(ctx)

	utils.CheckGinError(err, c)

	return conn
}

func (ah *DefaultAuthHandler) FullfillLogin(c *gin.Context, dbConn *pgxpool.Conn, user goth.User) {
	defer dbConn.Release()

	ctx := context.Background()

	queries := orm.New(dbConn)

	name := user.Name

	if len(name) == 0 {
		name = user.NickName
	}

	if len(name) == 0 {
		name = user.FirstName
	}

	if len(name) == 0 {
		name = "Unnamed User"
	}

	dbUser, err := queries.FindUserByProviderUserId(ctx, orm.FindUserByProviderUserIdParams{
		ProviderUserID: pgtype.Text{String: user.UserID, Valid: true},
		Provider:       pgtype.Text{String: user.Provider, Valid: true},
	})

	if err == nil {
		_ = queries.UpdateUser(ctx, orm.UpdateUserParams{
			DisplayName: pgtype.Text{String: name, Valid: true},
			ImagePath:   pgtype.Text{String: user.AvatarURL, Valid: true},
			ID:          dbUser.ID,
		})
	}

	if err == pgx.ErrNoRows {
		dbUser, err = queries.InsertUser(ctx, orm.InsertUserParams{
			ProviderUserID: pgtype.Text{String: user.UserID, Valid: true},
			Provider:       pgtype.Text{String: user.Provider, Valid: true},
			ImagePath:      pgtype.Text{String: user.AvatarURL, Valid: true},
			DisplayName:    pgtype.Text{String: user.Name, Valid: true},
			Email:          pgtype.Text{String: user.Email, Valid: true},
			Category:       CategoryCommon,
		})
		utils.CheckGinError(err, c)
	}

	ep := utils.NewDefaultEnvironmentProvider()

	if dbUser.Banned {
		frontendUrl := fmt.Sprintf("%s/login?banned=true", ep.GetBaseUrl())
		c.Redirect(http.StatusTemporaryRedirect, frontendUrl)
		return
	}

	authTokenManager := utils.NewJwtAuthTokenManager(ep)
	token, err := authTokenManager.CreateToken(dbUser)
	utils.CheckGinError(err, c)

	domain := ep.GetBackendDomain()
	maxAge := 24 * 60 * 60 * 1000
	bearerToken := "Bearer " + token

	c.SetCookie(TokenCookie, bearerToken, maxAge, "/", domain, true, false)

	frontendUrl := fmt.Sprintf("%s/login/callback", ep.GetBaseUrl())
	c.Redirect(http.StatusTemporaryRedirect, frontendUrl)
}

func (ah *DefaultAuthHandler) LogInUser(c *gin.Context) {
	gothic.GetProviderName = func(req *http.Request) (string, error) { return c.Param("provider"), nil }
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		conn := ah.getConn(c)
		ah.FullfillLogin(c, conn, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func (ah *DefaultAuthHandler) LogInUserCallback(c *gin.Context) {
	gothic.GetProviderName = func(req *http.Request) (string, error) { return c.Param("provider"), nil }
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	utils.CheckGinError(err, c)

	session := sessions.Default(c)
	session.Set("userID", user.UserID)
	session.Save()

	conn := ah.getConn(c)
	ah.FullfillLogin(c, conn, user)
}

func (ah *DefaultAuthHandler) LogOutUser(c *gin.Context) {
	err := gothic.Logout(c.Writer, c.Request)
	utils.CheckGinError(err, c)

	ep := utils.NewDefaultEnvironmentProvider()
	domain := ep.GetBackendDomain()
	c.SetCookie(TokenCookie, "", 0, "/", domain, true, false)

	c.Redirect(307, ep.GetBaseUrl())
}
