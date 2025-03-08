package auth

import (
	"backend/internal/utils"
	"backend/orm"
	"context"

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
	ctx := context.Background()

	queries := orm.New(dbConn)

	dbUser, err := queries.FindUserByEmail(ctx, user.Email)

	if err == pgx.ErrNoRows {
		dbUser, err = queries.InsertUser(ctx, orm.InsertUserParams{
			ImagePath:   pgtype.Text{String: user.AvatarURL},
			DisplayName: user.NickName,
			Category:    CategoryCommon,
			Email:       user.Email,
		})
		utils.CheckGinError(err, c)
	}

	ep := utils.NewDefaultEnvironmentProvider()
	authTokenManager := utils.NewJwtAuthTokenManager(ep)
	token, err := authTokenManager.CreateToken(dbUser)
	utils.CheckGinError(err, c)

	domain := ep.GetBackendDomain()
	maxAge := 24 * 60 * 60 * 1000

	c.SetCookie(TokenCookie, "Bearer "+token, maxAge, "/", domain, true, true)

	c.JSON(200, gin.H{
		"message": "successfully logged in",
	})
}

func (ah *DefaultAuthHandler) LogInUser(c *gin.Context) {
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		conn := ah.getConn(c)
		ah.FullfillLogin(c, conn, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func (ah *DefaultAuthHandler) LogInUserCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	utils.CheckGinError(err, c)

	conn := ah.getConn(c)
	ah.FullfillLogin(c, conn, user)
}

func (ah *DefaultAuthHandler) LogOutUser(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)

	c.JSON(200, gin.H{
		"message": "logout success",
	})
}
