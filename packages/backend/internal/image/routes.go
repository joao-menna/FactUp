package image

import (
	"backend/internal/middleware"
	"backend/internal/utils"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine) {
	r := g.Group("api/v1/image")

	cwd, err := os.Getwd()
	utils.CheckError(err)

	ih := NewDefaultImageHandler()

	r.POST("", middleware.AuthRequired(), ih.UploadImage)
	r.Static("", path.Join(cwd, "images"))
}
