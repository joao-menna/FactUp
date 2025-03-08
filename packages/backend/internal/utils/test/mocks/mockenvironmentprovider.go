package utilsmocks

import "backend/internal/utils"

type MockEnvironmentProviderProps struct {
	JwtSecretKey string
	Domain       string
	BaseUrl      string
}

type MockEnvironmentProvider struct {
	utils.EnvironmentProvider
	JwtSecretKey string
	Domain       string
	BaseUrl      string
}

func NewMockEnvironmentProvider(props MockEnvironmentProviderProps) *MockEnvironmentProvider {
	return &MockEnvironmentProvider{
		JwtSecretKey: props.JwtSecretKey,
		Domain:       props.Domain,
		BaseUrl:      props.BaseUrl,
	}
}

func (m *MockEnvironmentProvider) GetBackendJwtSecretKey() []byte {
	return []byte(m.JwtSecretKey)
}
