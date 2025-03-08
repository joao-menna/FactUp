package utils

import (
	"os"
	"strings"
)

type DefaultEnvironmentProvider struct {
	EnvironmentProvider
}

func NewDefaultEnvironmentProvider() *DefaultEnvironmentProvider {
	return &DefaultEnvironmentProvider{}
}

func (m *DefaultEnvironmentProvider) GetBackendJwtSecretKey() []byte {
	return []byte(os.Getenv("BACKEND_LOGIN_JWT_SECRET"))
}

func (p *DefaultEnvironmentProvider) GetBackendDomain() string {
	return os.Getenv("BACKEND_DOMAIN")
}

func (p *DefaultEnvironmentProvider) GetBackendPostgresConnectionUrl() string {
	return os.Getenv("BACKEND_POSTGRES_CONNECTION_URL")
}

func (p *DefaultEnvironmentProvider) GetProviderClientId(provider string) string {
	return os.Getenv("BACKEND_" + strings.ToUpper(provider) + "_CLIENT_ID")
}

func (p *DefaultEnvironmentProvider) GetProviderClientSecret(provider string) string {
	return os.Getenv("BACKEND_" + strings.ToUpper(provider) + "_CLIENT_SECRET")
}

func (p *DefaultEnvironmentProvider) GetProviderCallbackUrl(provider string) string {
	return os.Getenv("BACKEND_BASE_URL") + "/api/v1/auth/" + strings.ToLower(provider) + "/callback"
}
