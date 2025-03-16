package middleware

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(auth.TokenCookie)
		if err == http.ErrNoCookie {
			cookie = c.GetHeader(auth.TokenCookie)
		}

		if len(cookie) == 0 || !strings.HasPrefix(cookie, "Bearer ") {
			c.JSON(401, gin.H{
				"message": "user not logged in",
			})
			return
		}

		token := strings.Split(cookie, " ")[1]

		ep := utils.NewDefaultEnvironmentProvider()
		atm := utils.NewJwtAuthTokenManager(ep)

		parsed, err := atm.ValidateToken(token)
		utils.CheckGinError(err, c)

		c.Set(auth.UserID, parsed.UserID)
		c.Set(auth.Category, parsed.Category)

		c.Next()
	}
}
