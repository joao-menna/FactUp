package auth

import (
	"backend/internal/utils"
	"backend/orm"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func fullfillLogin(c *gin.Context, dbConn *pgxpool.Conn, user goth.User) {
	ctx := context.Background()

	queries := orm.New(dbConn)

	dbUser, err := queries.GetUserByEmail(ctx, user.Email)

	if err == pgx.ErrNoRows {
		dbUser, err = queries.InsertUser(ctx, orm.InsertUserParams{
			DisplayName: user.NickName,
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

	c.SetCookie("jwt", token, maxAge, "/", domain, true, true)
}

func loginHandler(c *gin.Context, dbPool *pgxpool.Pool) {
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		ctx := context.Background()

		conn, err := dbPool.Acquire(ctx)
		utils.CheckGinError(err, c)

		fullfillLogin(c, conn, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func loginCallbackHandler(c *gin.Context, dbPool *pgxpool.Pool) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn, err := dbPool.Acquire(ctx)
	utils.CheckGinError(err, c)

	fullfillLogin(c, conn, user)
}

func logoutHandler(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
}
