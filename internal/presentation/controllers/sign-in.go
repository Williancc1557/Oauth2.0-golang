package controllers

import (
	"encoding/json"
	dataProtocols "example/internal/data/protocols"
	"example/internal/domain/usecase"
	presentationProtocols "example/internal/presentation/protocols"
	"net/http"
)

type SignInController struct {
	GetAccountByEmail usecase.GetAccountByEmail
	Encrypter         dataProtocols.Encrypter
	ResetRefreshToken usecase.ResetRefreshToken
}

type SignInControllerBody struct {
	Email    string
	Password string
}

type SignInControllerResponse struct {
	RefreshToken string `json:"refreshToken"`
}

func (c *SignInController) Handle(r presentationProtocols.HttpRequest) (*presentationProtocols.HttpResponse, error) {
	var body SignInControllerBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	account, err := c.GetAccountByEmail.Get(body.Email)
	if err != nil {
		return nil, err
	}

	isCorrectPassword := c.Encrypter.Compare(body.Password, account.Password)
	if !isCorrectPassword {
		return nil, nil // mudar esses erros
	}

	newRefreshToken, err := c.ResetRefreshToken.Reset(account.Id)
	if err != nil {
		return nil, err
	}

	return &presentationProtocols.HttpResponse{
		Body: &SignInControllerResponse{
			RefreshToken: newRefreshToken,
		},
		StatusCode: http.StatusOK,
	}, nil
}
