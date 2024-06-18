package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/helpers"
	presentationProtocols "github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/go-playground/validator/v10"
)

type SignUpController struct {
	GetAccountByEmail usecase.GetAccountByEmail
	Validate          *validator.Validate
	AddAccount        usecase.AddAccount
	CreateAccessToken usecase.CreateAccessToken
}

func NewSignUpController(getAccountByEmail usecase.GetAccountByEmail) SignUpController {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return SignUpController{
		Validate:          validate,
		GetAccountByEmail: getAccountByEmail,
	}
}

type SignUpControllerBody struct {
	Email    string `validate:"email"`
	Password string `validate:"min=8,max=128"`
}

type SignUpControllerResponse struct {
	ExpiresIn    int    `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c *SignUpController) Handle(req presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	var body SignUpControllerBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "invalid body request",
		}, http.StatusBadRequest)
	}

	err = c.Validate.Struct(body)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: err.Error(),
		}, http.StatusUnprocessableEntity)
	}

	_, err = c.GetAccountByEmail.Get(body.Email)
	if err == nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "User already exists",
		}, http.StatusConflict)
	}

	addAccountInput := &usecase.AddAccountInput{
		Email:    body.Email,
		Password: body.Password,
	}
	account, err := c.AddAccount.Add(addAccountInput)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "An error ocurred while adding account",
		}, http.StatusBadRequest)
	}

	tokenData, err := c.CreateAccessToken.Create(account.Id)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "An error ocurred while creating access token",
		}, http.StatusBadRequest)
	}

	return helpers.CreateResponse(&SignUpControllerResponse{
		AccessToken:  tokenData.AccessToken,
		ExpiresIn:    tokenData.ExpiresIn,
		RefreshToken: account.RefreshToken,
	}, http.StatusOK)
}
