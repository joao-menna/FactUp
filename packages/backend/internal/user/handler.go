package user

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler interface {
	BanUser(c *gin.Context)
	CreateBot(c *gin.Context)
	ResetBotSecret(c *gin.Context)
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
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn := uh.getConn(c)

	queries := orm.New(conn)

	err = queries.BanUser(ctx, int32(id))
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message": "user banned",
	})
}

func (uh *DefaultUserHandler) CreateBot(c *gin.Context) {
	category, exists := c.Get(auth.Category)
	if !exists || category != auth.CategoryAdmin {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	err := c.ShouldBindJSON(&body)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn := uh.getConn(c)

	queries := orm.New(conn)

	secret, err := uuid.NewRandom()
	utils.CheckGinError(err, c)

	user, err := queries.InsertBotUser(ctx, orm.InsertBotUserParams{
		DisplayName: pgtype.Text{String: body.Name, Valid: true},
		Category:    auth.CategoryCommon,
	})
	utils.CheckGinError(err, c)

	bot, err := queries.InsertBot(ctx, orm.InsertBotParams{
		UserID: user.ID,
		Name:   body.Name,
		Secret: secret.String(),
	})
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message":  "bot created",
		"botToken": fmt.Sprintf("Bot %d_%s", bot.ID, bot.Secret),
	})
}

func (uh *DefaultUserHandler) ResetBotSecret(c *gin.Context) {
	category, exists := c.Get(auth.Category)
	if !exists || category != auth.CategoryAdmin {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	botIdStr := c.Param("id")
	botId, err := utils.ParseQueryId(botIdStr)
	utils.CheckGinError(err, c)

	ctx := context.Background()

	conn := uh.getConn(c)

	queries := orm.New(conn)

	secret, err := uuid.NewRandom()
	utils.CheckGinError(err, c)

	err = queries.UpdateBotSecret(ctx, orm.UpdateBotSecretParams{
		ID:     int32(botId),
		Secret: secret.String(),
	})
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message":  "bot created",
		"botToken": fmt.Sprintf("Bot %d_%s", botId, secret.String()),
	})
}
