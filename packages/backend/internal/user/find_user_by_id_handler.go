package user

import (
	"backend/internal/utils"
	"backend/orm"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func findUserByIdHandler(c *gin.Context, dbPool *pgxpool.Pool) {
	ctx := context.Background()

	conn, err := dbPool.Acquire(ctx)
	utils.CheckGinError(err, c)

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	utils.CheckGinError(err, c)

	queries := orm.New(conn)

	user, err := queries.GetUserById(ctx, int32(id))
	utils.CheckGinError(err, c)
	c.JSON(200, user)
}
