package controllers

import (
	"encoding/json"
	dataProtocols "example/internal/data/protocols"
	"example/internal/domain/usecase"
	presentationProtocols "example/internal/presentation/protocols"
	"example/internal/utils"
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

func (c *SignInController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	var body SignInControllerBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return utils.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid body request",
		}, http.StatusBadRequest)
	}

	account, err := c.GetAccountByEmail.Get(body.Email)
	if err != nil {
		return utils.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid credentials",
		}, http.StatusBadRequest)
	}

	isCorrectPassword := c.Encrypter.Compare(body.Password, account.Password)
	if !isCorrectPassword {
		return utils.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid credentials",
		}, http.StatusBadRequest)
	}

	newRefreshToken, err := c.ResetRefreshToken.Reset(account.Id)
	if err != nil {
		return utils.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while resetting refresh token",
		}, http.StatusBadRequest)
	}

	response := &SignInControllerResponse{
		RefreshToken: newRefreshToken,
	}

	return utils.CreateResponse(response, http.StatusOK)
}
