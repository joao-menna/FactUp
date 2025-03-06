package utilsmocks

import (
	"backend/internal/interfaces"
)

type MockEnvironmentProvider struct {
	interfaces.EnvironmentProvider
}

func NewMockEnvironmentProvider() *MockEnvironmentProvider {
	return &MockEnvironmentProvider{}
}

func (m *MockEnvironmentProvider) GetBackendJwtSecretKey() []byte {
	return []byte("secret_token_key")
}
