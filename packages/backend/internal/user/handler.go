package user

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler interface {
	BanUser(c *gin.Context)
}

type DefaultUserHandler struct {
	UserHandler
	dbPool *pgxpool.Pool
}

func NewDefaultUserHandler(dbPool *pgxpool.Pool) *DefaultUserHandler {
	return &DefaultUserHandler{
		dbPool: dbPool,
	}
}

func (uh *DefaultUserHandler) getConn(c *gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := uh.dbPool.Acquire(ctx)
	utils.CheckGinError(err, c)

	return conn
}

func (uh *DefaultUserHandler) BanUser(c *gin.Context) {
	category, exists := c.Get(auth.Category)
	if !exists || category != auth.CategoryAdmin {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	idStr := c.Param("userId")
	id, err := utils.ParseQueryId(idStr)

	ctx := context.Background()

	conn := uh.getConn(c)

	queries := orm.New(conn)

	err = queries.BanUser(ctx, int32(id))
	utils.CheckGinError(err, c)
	c.JSON(200, gin.H{
		"message": "user banned",
	})
}
