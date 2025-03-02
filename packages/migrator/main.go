package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		os.Getenv("MIGRATOR_POSTGRES_CONNECTION_URL"),
	)

	if err != nil {
		log.Panicln(err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Panicln(err)
	}

	log.Println("Database migrated successfully!")
}
