package interfaces

type EnvironmentProvider interface {
	GetBackendJwtSecretKey() []byte
	GetBackendDomain() string
	GetBackendPostgresConnectionUrl() string
	GetProviderClientId(provider string) string
	GetProviderClientSecret(provider string) string
	GetProviderCallbackUrl(provider string) string
}
