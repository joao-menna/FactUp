package main

import (
	"backend/internal/auth"
	"backend/internal/image"
	"backend/internal/interaction"
	"backend/internal/post"
	"backend/internal/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	ctx := context.Background()

	ep := utils.NewDefaultEnvironmentProvider()
	dbPool, err := pgxpool.New(ctx, ep.GetBackendPostgresConnectionUrl())
	utils.CheckError(err)
	defer dbPool.Close()

	r := gin.Default()

	image.Routes(r)
	auth.Routes(r, dbPool)
	post.Routes(r, dbPool)
	interaction.Routes(r, dbPool)

	err = r.Run(":8080")
	utils.CheckError(err)
}
