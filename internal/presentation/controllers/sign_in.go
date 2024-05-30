package controllers

import (
	"encoding/json"
	"net/http"

	dataProtocols "github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/helpers"
	presentationProtocols "github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
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
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid body request",
		}, http.StatusBadRequest)
	}

	account, err := c.GetAccountByEmail.Get(body.Email)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid credentials",
		}, http.StatusBadRequest)
	}

	isCorrectPassword := c.Encrypter.Compare(body.Password, account.Password)
	if !isCorrectPassword {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid credentials",
		}, http.StatusBadRequest)
	}

	newRefreshToken, err := c.ResetRefreshToken.Reset(account.Id)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "an error ocurred while resetting refresh token",
		}, http.StatusBadRequest)
	}

	response := &SignInControllerResponse{
		RefreshToken: newRefreshToken,
	}

	return helpers.CreateResponse(response, http.StatusOK)
}
