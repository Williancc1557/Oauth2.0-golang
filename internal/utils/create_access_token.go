package utils

import (
	"time"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/golang-jwt/jwt"
)

type CreateAccessTokenUtil struct {
	privateKey string
}

func NewCreateAccessTokenUtil(privateKey string) *CreateAccessTokenUtil {
	return &CreateAccessTokenUtil{
		privateKey: privateKey,
	}
}

func (b *CreateAccessTokenUtil) Create(userId string) (*usecase.CreateAccessTokenOutput, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	expiresIn := time.Now().Add(10 * time.Minute)
	claims["exp"] = expiresIn
	claims["authorized"] = true
	claims["sub"] = userId

	tokenString, err := token.SignedString([]byte(b.privateKey))
	if err != nil {
		return nil, err
	}

	return &usecase.CreateAccessTokenOutput{
		AccessToken: tokenString,
		ExpiresIn:   expiresIn.Second(),
	}, nil
}
