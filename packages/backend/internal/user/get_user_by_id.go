package user

import (
	"backend/internal/utils"
	"backend/orm"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByIdHandler(c *gin.Context, dbPool *pgxpool.Pool) {
	ctx := context.Background()

	conn, err := dbPool.Acquire(ctx)
	utils.CheckError(err)

	idStr, idExists := c.GetQuery("id")
	if !idExists {
		c.JSON(401, gin.H{
			"message": "ID not found in query",
		})
	}

	queries := orm.New(conn)

	if id, err := strconv.Atoi(idStr); err == nil {
		user, err := queries.GetUserById(ctx, int32(id))
		utils.CheckError(err)
		c.JSON(200, user)
	}
}
