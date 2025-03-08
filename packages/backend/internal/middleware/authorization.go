package middleware

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(auth.TokenCookie)
		utils.CheckGinError(err, c)

		token := strings.Split(cookie, " ")[1]

		ep := utils.NewDefaultEnvironmentProvider()
		atm := utils.NewJwtAuthTokenManager(ep)

		parsed, err := atm.ValidateToken(token)
		utils.CheckGinError(err, c)

		c.Set(auth.UserID, parsed.UserID)

		c.Next()
	}
}
