package main

import (
	"backend/internal/auth"
	"backend/internal/image"
	"backend/internal/interaction"
	"backend/internal/post"
	"backend/internal/user"
	"backend/internal/utils"
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	store := cookie.NewStore(ep.GetBackendJwtSecretKey())
	r.Use(sessions.Sessions("session", store))

	image.Routes(r, dbPool)
	auth.Routes(r, dbPool)
	post.Routes(r, dbPool)
	user.Routes(r, dbPool)
	interaction.Routes(r, dbPool)

	err = r.Run(":8080")
	utils.CheckError(err)
}
