package interaction

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserInteractionHandler interface {
	Add(c *gin.Context)
	Remove(c *gin.Context)
}

type DefaultUserInteractionHandler struct {
	UserInteractionHandler
	dbPool *pgxpool.Pool
}

func NewDefaultUserInteractionHandler(dbPool *pgxpool.Pool) *DefaultUserInteractionHandler {
	return &DefaultUserInteractionHandler{
		dbPool: dbPool,
	}
}

func (ah *DefaultUserInteractionHandler) getConn(c *gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := ah.dbPool.Acquire(ctx)

	utils.CheckGinError(err, c)

	return conn
}

func (uih *DefaultUserInteractionHandler) Add(c *gin.Context) {
	var body struct {
		PostID int `json:"postId"`
		Score  int `json:"score"`
	}

	userId, exists := c.Get(auth.UserID)
	if !exists {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	err := c.ShouldBindJSON(&body)
	utils.CheckGinError(err, c)

	conn := uih.getConn(c)

	queries := orm.New(conn)

	ctx := context.Background()

	_, err = queries.InsertUserInteraction(ctx, orm.InsertUserInteractionParams{
		PostID: int32(body.PostID),
		UserID: userId.(int32),
		Score:  int16(body.Score),
	})
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message": "post scored successfully",
	})
}

func (uih *DefaultUserInteractionHandler) Remove(c *gin.Context) {
	var body struct {
		PostID int `json:"postId"`
	}

	userId, exists := c.Get(auth.UserID)
	if !exists {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	err := c.ShouldBindJSON(&body)
	utils.CheckGinError(err, c)

	conn := uih.getConn(c)

	queries := orm.New(conn)

	ctx := context.Background()

	err = queries.DeleteUserInteraction(ctx, orm.DeleteUserInteractionParams{
		PostID: int32(body.PostID),
		UserID: userId.(int32),
	})
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message": "post scored successfully",
	})
}
