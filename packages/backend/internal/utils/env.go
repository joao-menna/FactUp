package utils

import "os"

func GetBackendDomain() string {
	return os.Getenv("BACKEND_DOMAIN")
}

func GetBackendPostgresConnectionUrl() string {
	return os.Getenv("BACKEND_POSTGRES_CONNECTION_URL")
}
