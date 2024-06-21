package utils

import (
	"os"
	"time"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/golang-jwt/jwt"
)

type CreateAccessTokenUtil struct{}

func NewCreateAccessTokenUtil() *CreateAccessTokenUtil {
	return &CreateAccessTokenUtil{}
}

func (b *CreateAccessTokenUtil) Create(userId string) (*usecase.CreateAccessTokenOutput, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	expiresIn := time.Now().Add(time.Minute * 5).Unix()
	claims["exp"] = expiresIn
	claims["authorized"] = true
	claims["sub"] = userId

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	return &usecase.CreateAccessTokenOutput{
		AccessToken: tokenString,
		ExpiresIn:   int(expiresIn),
	}, nil
}
