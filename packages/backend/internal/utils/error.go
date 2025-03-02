package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func CheckGinError(err error, c *gin.Context) {
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		log.Panicln(err)
	}
}
