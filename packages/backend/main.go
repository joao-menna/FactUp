package main

import (
	"backend/internal/user"
	"backend/internal/utils"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, os.Getenv("BACKEND_POSTGRES_CONNECTION_URL"))
	utils.CheckError(err)
	defer dbPool.Close()

	r := gin.Default()

	user.Routes(r, dbPool)

	err = r.Run(":8080")
	utils.CheckError(err)
}
