package image

import (
	"backend/internal/utils"
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "golang.org/x/image/webp"
)

type ImageHandler interface {
	UploadImage(c *gin.Context)
}

type DefaultImageHandler struct {
	ImageHandler
}

func NewDefaultImageHandler() *DefaultImageHandler {
	return &DefaultImageHandler{}
}

func (ih *DefaultImageHandler) UploadImage(c *gin.Context) {
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

	c.JSON(200, gin.H{
		"message":   "image received",
		"imagePath": imagePath,
	})
}
