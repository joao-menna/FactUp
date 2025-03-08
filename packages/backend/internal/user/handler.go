package user

import (
	"backend/internal/utils"
	"backend/orm"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler interface {
	MakeUserAdmin(c *gin.Context)
}

type DefaultUserHandler struct {
	UserHandler
	dbPool *pgxpool.Pool
}

func (uh *DefaultUserHandler) getConn(c *gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := uh.dbPool.Acquire(ctx)
	utils.CheckGinError(err, c)

	return conn
}

func (uh *DefaultUserHandler) MakeUserAdmin(c *gin.Context) {
	conn := uh.getConn(c)

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	utils.CheckGinError(err, c)

	queries := orm.New(conn)

	ctx := context.Background()

	user, err := queries.FindUserById(ctx, int32(id))
	utils.CheckGinError(err, c)
	c.JSON(200, user)
}
