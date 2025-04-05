package main

import (
	"backend/internal/auth"
	"backend/internal/image"
	"backend/internal/interaction"
	"backend/internal/post"
	"backend/internal/user"
	"backend/internal/utils"
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	ep := utils.NewDefaultEnvironmentProvider()
	dbPool, err := pgxpool.New(ctx, ep.GetBackendPostgresConnectionUrl())
	utils.CheckError(err)
	defer dbPool.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:5173", "https://factup.me"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
