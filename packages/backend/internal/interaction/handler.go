package interaction

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserInteractionHandler interface {
	CheckPostShouldDelete(postId int)
	GetByPostId(c *gin.Context)
	GetForMultipleByUserId(c *gin.Context)
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

func (uih *DefaultUserInteractionHandler) getConn(c ...*gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := uih.dbPool.Acquire(ctx)

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

func (uih *DefaultUserInteractionHandler) GetForMultipleByUserId(c *gin.Context) {
	userId, exists := c.Get(auth.UserID)
	if !exists {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	postIdsStr := c.QueryArray("postId")

	postIds, err := utils.StringSliceToIntSlice(postIdsStr)
	utils.CheckGinError(err, c)

	postIds32 := utils.IntSliceToInt32Slice(postIds)

	conn := uih.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	ctx := context.Background()

	uis, err := queries.FindInteractionByUserIdAndMultiplePostIds(ctx, orm.FindInteractionByUserIdAndMultiplePostIdsParams{
		UserID:  userId.(int32),
		Column2: postIds32,
	})
	utils.CheckGinError(err, c)

	c.JSON(200, uis)
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

	err = utils.ValidateScore(body.Score)
	utils.CheckGinError(err, c)

	conn := uih.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	ctx := context.Background()

	_, err = queries.FindInteractionByUserIdAndPostId(ctx, orm.FindInteractionByUserIdAndPostIdParams{
		PostID: int32(body.PostID),
		UserID: userId.(int32),
	})

	if errors.Is(err, pgx.ErrNoRows) {
		_, err = queries.InsertUserInteraction(ctx, orm.InsertUserInteractionParams{
			PostID: int32(body.PostID),
			UserID: userId.(int32),
			Score:  int16(body.Score),
		})
		utils.CheckGinError(err, c)
	} else {
		err := queries.UpdateUserInteraction(ctx, orm.UpdateUserInteractionParams{
			PostID: int32(body.PostID),
			UserID: userId.(int32),
			Score:  int16(body.Score),
		})
		utils.CheckGinError(err, c)
	}

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

	postIdStr := c.Query("postId")
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
