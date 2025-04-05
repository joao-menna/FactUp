package interaction

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserInteractionHandler interface {
	CheckPostShouldDelete(postId int)
	GetByPostId(c *gin.Context)
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

func (ah *DefaultUserInteractionHandler) getConn(c ...*gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := ah.dbPool.Acquire(ctx)

	if len(c) > 0 {
		utils.CheckGinError(err, c[0])
	}

	return conn
}

func (uih *DefaultUserInteractionHandler) CheckPostShouldDelete(postId int) {
	ctx := context.Background()

	conn := uih.getConn()
	defer conn.Release()

	queries := orm.New(conn)

	score, err := queries.GetInteractionScoreByPostId(ctx, int32(postId))

	utils.CheckError(err)

	err = utils.CheckIfScoreShouldDeletePost(int(score))

	if err != utils.ErrMaxNegativeScoreReached {
		return
	}

	err = queries.DeletePostById(ctx, int32(postId))
	if err != nil {
		log.Panicln(err)
	}
}

func (uih *DefaultUserInteractionHandler) GetByPostId(c *gin.Context) {
	postIdStr := c.Param("postId")
	postId, err := utils.ParseQueryId(postIdStr)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn := uih.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	score, err := queries.GetInteractionScoreByPostId(ctx, int32(postId))

	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"score": score,
	})
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
	defer conn.Release()

	queries := orm.New(conn)

	ctx := context.Background()

	_, err = queries.InsertUserInteraction(ctx, orm.InsertUserInteractionParams{
		PostID: int32(body.PostID),
		UserID: userId.(int32),
		Score:  int16(body.Score),
	})
	utils.CheckGinError(err, c)

	go uih.CheckPostShouldDelete(body.PostID)

	c.JSON(200, gin.H{
		"message": "post scored successfully",
	})
}

func (uih *DefaultUserInteractionHandler) Remove(c *gin.Context) {
	userId, exists := c.Get(auth.UserID)
	if !exists {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	postIdStr := c.Param("postId")
	postId, err := utils.ParseQueryId(postIdStr)
	utils.CheckGinError(err, c)

	conn := uih.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	ctx := context.Background()

	err = queries.DeleteUserInteraction(ctx, orm.DeleteUserInteractionParams{
		PostID: int32(postId),
		UserID: userId.(int32),
	})
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message": "post scored successfully",
	})
}
