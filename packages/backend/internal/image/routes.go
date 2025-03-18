package image

import (
	"backend/internal/middleware"
	"backend/internal/utils"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	r := g.Group("api/v1/image")

	cwd, err := os.Getwd()
	utils.CheckError(err)

	ih := NewDefaultImageHandler()

	r.POST("", middleware.AuthRequired(dbPool), ih.UploadImage)
	r.Static("images", path.Join(cwd, "images"))
}
