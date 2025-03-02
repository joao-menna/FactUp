package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	r := g.Group("api/v1/user")

	r.GET(":id", func(c *gin.Context) { findUserByIdHandler(c, dbPool) })
}
