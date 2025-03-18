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
	ctx := context.Background()

	queries := orm.New(dbConn)

	dbUser, err := queries.FindUserByEmail(ctx, pgtype.Text{String: user.Email, Valid: true})

	if err == pgx.ErrNoRows {
		dbUser, err = queries.InsertUser(ctx, orm.InsertUserParams{
			ImagePath:   pgtype.Text{String: user.AvatarURL, Valid: true},
			DisplayName: pgtype.Text{String: user.NickName, Valid: true},
			Category:    CategoryCommon,
			Email:       pgtype.Text{String: user.Email, Valid: true},
		})
		utils.CheckGinError(err, c)
	}

	if dbUser.Banned {
		c.JSON(401, gin.H{
			"message": "you are banned",
		})
		return
	}

	ep := utils.NewDefaultEnvironmentProvider()
	authTokenManager := utils.NewJwtAuthTokenManager(ep)
	token, err := authTokenManager.CreateToken(dbUser)
	utils.CheckGinError(err, c)

	domain := ep.GetBackendDomain()
	maxAge := 24 * 60 * 60 * 1000
	bearerToken := "Bearer " + token

	c.SetCookie(TokenCookie, bearerToken, maxAge, "/", domain, true, true)

	if c.Request.Header.Get("Host") == "" {
		c.JSON(200, gin.H{
			"message": "successfully logged in",
			"token":   bearerToken,
		})
	} else {
		frontendUrl := fmt.Sprintf("%s/?token=%s", ep.GetBaseUrl(), bearerToken)
		c.Redirect(http.StatusTemporaryRedirect, frontendUrl)
	}
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

	c.JSON(200, gin.H{
		"message": "logout success",
	})
}
