package utils

import (
	"backend/orm"
	utilsmocks "backend/test/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	user := orm.User{ID: 123}

	ep := utilsmocks.NewMockEnvironmentProvider()
	m := NewJwtAuthTokenManager(ep)

	token, err := m.CreateToken(user)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	parsed, err := jwt.ParseWithClaims(token, &TokenContent{}, func(t *jwt.Token) (any, error) {
		return ep.GetBackendJwtSecretKey(), nil
	})

	require.NoError(t, err)

	claims, ok := parsed.Claims.(*TokenContent)

	require.True(t, ok)
	require.Equal(t, int32(123), claims.UserID)
}
