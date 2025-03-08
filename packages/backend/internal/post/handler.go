package post

import (
	"backend/internal/utils"
	"backend/orm"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostHandler interface {
	FindRandom(c *gin.Context)
	FindById(c *gin.Context)
	FindAllByUser(c *gin.Context)
	InsertPost(c *gin.Context)
	DeletePostById(c *gin.Context)
}

type DefaultPostHandler struct {
	PostHandler
	dbPool *pgxpool.Pool
}

func NewDefaultPostHandler(dbPool *pgxpool.Pool) *DefaultPostHandler {
	return &DefaultPostHandler{
		dbPool: dbPool,
	}
}

func (ph *DefaultPostHandler) getConn(c *gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := ph.dbPool.Acquire(ctx)

	utils.CheckGinError(err, c)

	return conn
}

func (ph *DefaultPostHandler) FindRandom(c *gin.Context) {
	conn := ph.getConn(c)

	postType := c.Query("type")
	err := utils.ValidatePostType(postType)
	utils.CheckGinError(err, c)

	limitStr := c.Query("limit")
	limit, err := utils.ParsePostLimit(limitStr)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	queries := orm.New(conn)

	posts, err := queries.FindRandomPosts(ctx, orm.FindRandomPostsParams{
		Type:  postType,
		Limit: int32(limit),
	})

	utils.CheckGinError(err, c)

	c.JSON(200, posts)
}

func (ph *DefaultPostHandler) FindById(c *gin.Context) {
	conn := ph.getConn(c)

	postIdStr := c.Param("id")
	postId, err := utils.ParseQueryId(postIdStr)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	queries := orm.New(conn)

	post, err := queries.FindPostById(ctx, int32(postId))

	utils.CheckGinError(err, c)

	c.JSON(200, post)
}

func (ph *DefaultPostHandler) FindAllByUser(c *gin.Context) {
	conn := ph.getConn(c)

	userIdStr := c.Param("userId")
	userId, err := utils.ParseQueryId(userIdStr)
	utils.CheckGinError(err, c)

	limitStr := c.Query("limit")
	limit, err := utils.ParsePostLimit(limitStr)
	utils.CheckGinError(err, c)

	pageStr := c.Query("page")
	page, err := utils.ParseQueryId(pageStr)
	utils.CheckGinError(err, c)

	offset := utils.GetPostOffset(limit, page)

	ctx := context.Background()

	queries := orm.New(conn)

	post, err := queries.FindPostsByUserId(ctx, orm.FindPostsByUserIdParams{
		UserID: int32(userId),
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	utils.CheckGinError(err, c)

	c.JSON(200, post)
}

func (ph *DefaultPostHandler) InsertPost(c *gin.Context) {
	// TODO: InsertPost
}

func (ph *DefaultPostHandler) DeletePostById(c *gin.Context) {
	// TODO: DeletePostById
}
