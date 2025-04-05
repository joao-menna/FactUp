package post

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
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
	postType := c.Query("type")
	err := utils.ValidatePostType(postType)
	utils.CheckGinError(err, c)

	limitStr := c.Query("limit")
	limit, err := utils.ParsePostLimit(limitStr)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn := ph.getConn(c)

	queries := orm.New(conn)

	posts, err := queries.FindRandomPosts(ctx, orm.FindRandomPostsParams{
		Type:  postType,
		Limit: int32(limit),
	})

	utils.CheckGinError(err, c)

	c.JSON(200, posts)
}

func (ph *DefaultPostHandler) FindById(c *gin.Context) {
	postIdStr := c.Param("id")
	postId, err := utils.ParseQueryId(postIdStr)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn := ph.getConn(c)

	queries := orm.New(conn)

	post, err := queries.FindPostById(ctx, int32(postId))

	utils.CheckGinError(err, c)

	c.JSON(200, post)
}

func (ph *DefaultPostHandler) FindAllByUser(c *gin.Context) {
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

	conn := ph.getConn(c)

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
	var body struct {
		Type      string `json:"type"`
		Body      string `json:"body"`
		Source    string `json:"source"`
		ImagePath string `json:"imagePath"`
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

	ctx := context.Background()

	conn := ph.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	count, err := queries.GetPostedCountByDay(ctx, userId.(int32))
	utils.CheckGinError(err, c)

	err = utils.CheckPostMaxCountByDay(int(count))
	if err != nil {
		if err == utils.ErrMaxPostForToday {
			c.JSON(401, gin.H{
				"message": "reached maximum post count for this account today",
			})
			return
		}
		utils.CheckGinError(err, c)
	}

	post, err := queries.InsertPost(ctx, orm.InsertPostParams{
		Type:      body.Type,
		UserID:    userId.(int32),
		Body:      body.Body,
		Source:    pgtype.Text{String: body.Source, Valid: true},
		ImagePath: pgtype.Text{String: body.ImagePath, Valid: true},
	})

	utils.CheckGinError(err, c)

	c.JSON(200, post)
}

func (ph *DefaultPostHandler) DeletePostById(c *gin.Context) {
	postIdStr := c.Param("postId")
	postId, err := utils.ParseQueryId(postIdStr)
	utils.CheckGinError(err, c)

	userId, exists := c.Get(auth.UserID)
	category, _ := c.Get(auth.Category)
	if !exists {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	ctx := context.Background()

	conn := ph.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	post, err := queries.FindPostById(ctx, int32(postId))
	utils.CheckGinError(err, c)

	if userId != post.UserID || category.(string) != auth.CategoryAdmin {
		c.JSON(401, gin.H{
			"message": "this post is not yours",
		})
		return
	}

	err = queries.DeletePostById(ctx, int32(postId))

	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message": "post deleted successfully",
	})
}
