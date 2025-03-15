package interaction

import (
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	r := g.Group("/api/v1/interaction")

	r.Use(middleware.AuthRequired())

	uih := NewDefaultUserInteractionHandler(dbPool)

	r.GET(":postId", uih.GetByPostId)
	r.POST("", uih.Add)
	r.DELETE("", uih.Remove)
}
