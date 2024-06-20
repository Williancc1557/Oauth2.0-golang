package controllers

import (
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/helpers"
	presentationProtocols "github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
)

type RefreshTokenController struct {
	GetUserByRefreshToken usecase.GetUserByRefreshToken
	CreateAccessToken     usecase.CreateAccessToken
}

type RefreshTokenControllerOutput struct {
	ExpiresIn   int    `json:"expiresIn"`
	AccessToken string `json:"accessToken"`
}

func NewRefreshTokenController() *RefreshTokenController {
	return &RefreshTokenController{}
}

func (c *RefreshTokenController) Handle(r presentationProtocols.HttpRequest) *presentationProtocols.HttpResponse {
	refreshToken := r.Header.Get("refreshtoken")

	account, err := c.GetUserByRefreshToken.Get(refreshToken)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "An invalid refresh token was provided",
		}, http.StatusUnprocessableEntity)
	}

	accessTokenData, err := c.CreateAccessToken.Create(account.Id)
	if err != nil {
		return helpers.CreateResponse(&presentationProtocols.ErrorResponse{
			Error: "An error ocurred while creating access token",
		}, http.StatusUnprocessableEntity)
	}

	return helpers.CreateResponse(&RefreshTokenControllerOutput{
		ExpiresIn:   accessTokenData.ExpiresIn,
		AccessToken: accessTokenData.AccessToken,
	}, http.StatusOK)
}
