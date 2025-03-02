package utils

import (
	"os"
	"strings"
)

func GetProviderClientId(provider string) string {
	return os.Getenv("BACKEND_" + strings.ToUpper(provider) + "_CLIENT_ID")
}

func GetProviderClientSecret(provider string) string {
	return os.Getenv("BACKEND_" + strings.ToUpper(provider) + "_CLIENT_SECRET")
}

func GetProviderCallbackUrl(provider string) string {
	return os.Getenv("BACKEND_BASE_URL") + "/api/v1/auth/" + strings.ToLower(provider) + "/callback"
}
