package image

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"backend/orm"
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "golang.org/x/image/webp"
)

type ImageHandler interface {
	UploadImage(c *gin.Context)
}

type DefaultImageHandler struct {
	ImageHandler
	dbPool *pgxpool.Pool
}

func NewDefaultImageHandler(dbPool *pgxpool.Pool) *DefaultImageHandler {
	return &DefaultImageHandler{
		dbPool: dbPool,
	}
}

func (ih *DefaultImageHandler) getConn(c *gin.Context) *pgxpool.Conn {
	ctx := context.Background()

	conn, err := ih.dbPool.Acquire(ctx)

	utils.CheckGinError(err, c)

	return conn
}

func (ih *DefaultImageHandler) UploadImage(c *gin.Context) {
	userId, exists := c.Get(auth.UserID)
	if !exists {
		c.JSON(401, gin.H{
			"message": "user not logged in",
		})
		return
	}

	ctx := context.Background()

	conn := ih.getConn(c)
	defer conn.Release()

	queries := orm.New(conn)

	images, err := queries.GetImagePostedByUserId(ctx, userId.(int32))
	utils.CheckGinError(err, c)

	for _, i := range images {
		imagePath := fmt.Sprintf("images/%s.webp", i.ImagePath)
		postsWithImage, err := queries.FindPostsByImagePath(ctx, pgtype.Text{String: imagePath, Valid: true})
		utils.CheckGinError(err, c)

		if len(postsWithImage) == 0 {
			err := queries.DeleteImageById(ctx, i.ID)
			utils.CheckGinError(err, c)

			err = os.Remove(imagePath)
			if err != nil && !os.IsNotExist(err) {
				utils.CheckGinError(err, c)
			}
		}
	}

	totalInDay, err := queries.GetImagePostedInDayByUserId(ctx, userId.(int32))
	utils.CheckGinError(err, c)

	err = utils.CheckPostMaxCountByDay(int(totalInDay))
	utils.CheckGinError(err, c)

	header, err := c.FormFile("image")
	utils.CheckGinError(err, c)

	file, err := header.Open()
	utils.CheckGinError(err, c)
	defer file.Close()

	img, _, err := image.Decode(file)
	utils.CheckGinError(err, c)

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Lossless: true})
	utils.CheckGinError(err, c)

	uuid, err := uuid.NewRandom()
	utils.CheckGinError(err, c)

	imagePath := "images/" + uuid.String() + ".webp"

	err = os.WriteFile(imagePath, buf.Bytes(), 0666)
	utils.CheckGinError(err, c)

	_, err = queries.InsertImage(ctx, orm.InsertImageParams{
		UserID:    userId.(int32),
		ImagePath: uuid.String(),
	})
	utils.CheckGinError(err, c)

	c.JSON(200, gin.H{
		"message":   "image received",
		"imagePath": imagePath,
	})
}
