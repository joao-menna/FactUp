package user

import (
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	r := g.Group("api/v1/user")

	uh := NewDefaultUserHandler(dbPool)

	r.DELETE("ban/:userId", middleware.AuthRequired(), uh.BanUser)
}
