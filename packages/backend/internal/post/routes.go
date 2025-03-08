package post

import (
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	r := g.Group("/api/v1/post")

	ph := NewDefaultPostHandler(dbPool)

	r.GET("single/:id", ph.FindById)
	r.GET("multiple/random", ph.FindRandom)
	r.GET("multiple/user/:userId", ph.FindAllByUser)
	r.POST("", middleware.AuthRequired(), ph.InsertPost)
	r.DELETE(":id", middleware.AuthRequired(), ph.DeletePostById)
}
