package controllers

import (
	"bytes"
	"encoding/json"
	dataProtocols "example/internal/data/protocols"
	"example/internal/domain/usecase"
	presentationProtocols "example/internal/presentation/protocols"
	"example/internal/utils"
	"io"
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

func (c *SignInController) Handle(r presentationProtocols.HttpRequest) (*presentationProtocols.HttpResponse, *presentationProtocols.ErrorResponse) {
	var body SignInControllerBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, utils.HandleError("invalid body request", http.StatusBadRequest)
	}

	account, err := c.GetAccountByEmail.Get(body.Email)
	if err != nil {
		return nil, utils.HandleError("invalid credentials", http.StatusBadRequest)
	}

	isCorrectPassword := c.Encrypter.Compare(body.Password, account.Password)
	if !isCorrectPassword {
		return nil, utils.HandleError("invalid credentials", http.StatusBadRequest)
	}

	newRefreshToken, err := c.ResetRefreshToken.Reset(account.Id)
	if err != nil {
		return nil, utils.HandleError("an error ocurred while resetting refresh token", http.StatusBadRequest)
	}

	refreshTokenResponse, err := json.Marshal(&SignInControllerResponse{
		RefreshToken: newRefreshToken,
	})
	if err != nil {
		return nil, utils.HandleError("an error ocurred while marshaling response", http.StatusBadRequest)
	}

	return &presentationProtocols.HttpResponse{
		Body:       io.NopCloser(bytes.NewReader(refreshTokenResponse)),
		StatusCode: http.StatusOK,
	}, nil
}
