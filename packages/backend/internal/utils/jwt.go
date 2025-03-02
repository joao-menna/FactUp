package utils

import (
	"backend/orm"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtContent struct {
	UserID int32
}

func getKey() []byte {
	return []byte(os.Getenv("BACKEND_LOGIN_JWT_SECRET"))
}

func CreateLoginJwt(user orm.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": user.ID,
	})

	return token.SignedString(getKey())
}

func ValidateLoginJwt(token string) (*JwtContent, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return getKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok {
		return &JwtContent{
			UserID: claims["userId"].(int32),
		}, nil
	}

	return nil, errors.New("could not parse jwt claims")
}
