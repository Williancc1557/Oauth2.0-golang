package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupRefreshTokenMocks(t *testing.T) (*controllers.RefreshTokenController, *mocks.MockGetAccountByRefreshToken, *mocks.MockCreateAccessToken, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	getAccountByRefreshToken := mocks.NewMockGetAccountByRefreshToken(ctrl)
	createAccessToken := mocks.NewMockCreateAccessToken(ctrl)
	sut := controllers.NewRefreshTokenController(getAccountByRefreshToken, createAccessToken)

	return sut, getAccountByRefreshToken, createAccessToken, ctrl
}

func CreateHttpRequest(t *testing.T) protocols.HttpRequest {
	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(&controllers.SignInControllerBody{})
	require.NoError(t, err)

	header := http.Header{}
	header.Add("refreshtoken", "fake-refresh-token")

	return protocols.HttpRequest{
		Body:   io.NopCloser(&requestBody),
		Header: header,
	}
}

func TestRefreshTokenController(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, getAccountByRefreshToken, createAccessToken, ctrl := setupRefreshTokenMocks(t)
		defer ctrl.Finish()

		account := &models.AccountModel{
			Id:           "fake-account-id",
			Email:        "fake-account-email@example.com",
			Password:     "fake-account-password",
			RefreshToken: refreshToken,
		}

		accessTokenData := &usecase.CreateAccessTokenOutput{
			AccessToken: "fake-access-token",
			ExpiresIn:   123,
		}

		getAccountByRefreshToken.EXPECT().Get("fake-refresh-token").Return(account, nil)
		createAccessToken.EXPECT().Create("fake-account-id").Return(accessTokenData, nil)

		requestData := CreateHttpRequest(t)
		response := sut.Handle(requestData)

		require.Equal(t, response.StatusCode, http.StatusOK)

		var responseBody controllers.RefreshTokenControllerResponse
		err := json.NewDecoder(response.Body).Decode(&responseBody)
		require.NoError(t, err)

		correctResponse := &controllers.RefreshTokenControllerResponse{
			AccessToken: accessTokenData.AccessToken,
			ExpiresIn:   accessTokenData.ExpiresIn,
		}
		require.Equal(t, &responseBody, correctResponse)

	})
}
