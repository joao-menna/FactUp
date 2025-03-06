package utils

import (
	"backend/internal/interfaces"
	"backend/orm"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenContent struct {
	UserID int32 `json:"userId"`
	jwt.RegisteredClaims
}

type AuthTokenManager interface {
	CreateToken(user orm.User) (string, error)
	ValidateToken(token string) (*TokenContent, error)
}

type JwtAuthTokenManager struct {
	AuthTokenManager
	environmentProvider interfaces.EnvironmentProvider
}

func NewJwtAuthTokenManager(environmentProvider interfaces.EnvironmentProvider) *JwtAuthTokenManager {
	return &JwtAuthTokenManager{
		environmentProvider: environmentProvider,
	}
}

func (m *JwtAuthTokenManager) CreateToken(user orm.User) (string, error) {
	exp := time.Now().Add(time.Duration(24) * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenContent{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp.Unix(), 0)),
		},
	})

	return token.SignedString(m.environmentProvider.GetBackendJwtSecretKey())
}

func (m *JwtAuthTokenManager) ValidateToken(token string) (*TokenContent, error) {
	parsed, err := jwt.ParseWithClaims(token, &TokenContent{}, func(t *jwt.Token) (any, error) {
		return m.environmentProvider.GetBackendJwtSecretKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsed.Claims.(*TokenContent); ok {
		return claims, nil
	}

	return nil, errors.New("could not parse jwt claims")
}
