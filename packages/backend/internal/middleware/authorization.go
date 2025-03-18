package middleware

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRequired(dbPool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(auth.TokenCookie)
		if err == http.ErrNoCookie {
			cookie = c.Request.Header.Get(auth.TokenCookie)
		}

		if len(cookie) == 0 {
			c.JSON(401, gin.H{
				"message": "user not logged in",
			})
			return
		}

		if !strings.HasPrefix(cookie, "Bearer ") && !strings.HasPrefix(cookie, "Bot ") {
			c.JSON(401, gin.H{
				"message": "invalid token",
			})
			return
		}

		token := strings.Split(cookie, " ")[1]

		if strings.HasPrefix(cookie, "Bearer ") {
			validateUserToken(token, c)
			return
		}

		if strings.HasPrefix(cookie, "Bot ") {
			validateBotToken(token, c, dbPool)
			return
		}
	}
}

func validateUserToken(token string, c *gin.Context) {
	ep := utils.NewDefaultEnvironmentProvider()
	atm := utils.NewJwtAuthTokenManager(ep)

	parsed, err := atm.ValidateToken(token)
	utils.CheckGinError(err, c)

	c.Set(auth.UserID, parsed.UserID)
	c.Set(auth.Category, parsed.Category)

	c.Next()
}

func validateBotToken(token string, c *gin.Context, dbPool *pgxpool.Pool) {
	ctx := context.Background()

	conn, err := dbPool.Acquire(ctx)
	utils.CheckGinError(err, c)

	tokenSplit := strings.Split(token, "_")
	botIdStr := tokenSplit[0]
	botSecret := tokenSplit[1]
	botId, err := strconv.Atoi(botIdStr)
	utils.CheckGinError(err, c)

	queries := orm.New(conn)

	bot, err := queries.FindBotById(ctx, int32(botId))
	utils.CheckGinError(err, c)

	if botSecret != bot.Secret {
		c.JSON(401, gin.H{
			"message": "invalid bot secret",
		})
		return
	}

	user, err := queries.FindUserById(ctx, int32(bot.UserID))
	utils.CheckGinError(err, c)

	c.Set(auth.UserID, bot.UserID)
	c.Set(auth.Category, user.Category)

	c.Next()
}
