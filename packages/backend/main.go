package main

import (
	"backend/internal/auth"
	"backend/internal/utils"
	"context"

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

	auth.Routes(r, dbPool)

	err = r.Run(":8080")
	utils.CheckError(err)
}
