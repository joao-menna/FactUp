package utils_test

import (
	"backend/internal/utils"
	utilsmocks "backend/internal/utils/test/mocks"
	"backend/orm"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	user := orm.User{ID: 123}

	ep := utilsmocks.NewMockEnvironmentProvider(utilsmocks.MockEnvironmentProviderProps{
		JwtSecretKey: "test_secret_key",
	})
	m := utils.NewJwtAuthTokenManager(ep)

	token, err := m.CreateToken(user)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	parsed, err := jwt.ParseWithClaims(token, &utils.TokenContent{}, func(t *jwt.Token) (any, error) {
		return ep.GetBackendJwtSecretKey(), nil
	})

	require.NoError(t, err)

	claims, ok := parsed.Claims.(*utils.TokenContent)

	require.True(t, ok)
	require.Equal(t, int32(123), claims.UserID)
}

func TestValidateToken(t *testing.T) {
	user := orm.User{ID: 123}

	ep := utilsmocks.NewMockEnvironmentProvider(utilsmocks.MockEnvironmentProviderProps{
		JwtSecretKey: "test_secret_key",
	})
	m := utils.NewJwtAuthTokenManager(ep)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &utils.TokenContent{
		UserID: user.ID,
	})

	tokenStr, err := token.SignedString(ep.GetBackendJwtSecretKey())

	require.NoError(t, err)
	require.NotEmpty(t, token)

	parsed, err := m.ValidateToken(tokenStr)

	require.NoError(t, err)
	require.Equal(t, int32(123), parsed.UserID)
}

func TestTokenFlow(t *testing.T) {
	user := orm.User{ID: 123}

	ep := utilsmocks.NewMockEnvironmentProvider(utilsmocks.MockEnvironmentProviderProps{
		JwtSecretKey: "test_secret_key",
	})
	m := utils.NewJwtAuthTokenManager(ep)

	tokenStr, err := m.CreateToken(user)

	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)

	parsedToken, err := m.ValidateToken(tokenStr)

	require.NoError(t, err)
	require.Equal(t, user.ID, parsedToken.UserID)
}
