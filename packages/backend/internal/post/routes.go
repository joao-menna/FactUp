package post

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	r := g.Group("/api/v1/post")

	r.GET("multiple/random")
	r.GET("single/:id")
	r.GET("user/:userId")
}
