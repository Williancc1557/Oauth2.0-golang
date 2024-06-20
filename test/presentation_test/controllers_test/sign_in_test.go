package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func setupMocks(t *testing.T) (*controllers.SignInController, *mocks.MockEncrypter, *mocks.MockGetAccountByEmail, *mocks.MockResetRefreshToken, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockEncrypter := mocks.NewMockEncrypter(ctrl)
	mockGetAccountByEmail := mocks.NewMockGetAccountByEmail(ctrl)
	mockResetRefreshToken := mocks.NewMockResetRefreshToken(ctrl)

	signInController := controllers.NewSignInController(
		mockGetAccountByEmail,
		mockEncrypter,
		mockResetRefreshToken,
	)

	return signInController, mockEncrypter, mockGetAccountByEmail, mockResetRefreshToken, ctrl
}

func createHttpRequest(t *testing.T, email, password string) *protocols.HttpRequest {
	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(&controllers.SignInControllerBody{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)

	return &protocols.HttpRequest{
		Body:   io.NopCloser(&requestBody),
		Header: nil,
	}
}

func TestSignInController(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		signInController, mockEncrypter, mockGetAccountByEmail, mockResetRefreshToken, ctrl := setupMocks(t)
		defer ctrl.Finish()

		account := &models.AccountModel{
			Id:           "fake-account-id",
			Email:        "fake-account-email@example.com",
			Password:     "fake-account-password",
			RefreshToken: refreshToken,
		}

		mockGetAccountByEmail.EXPECT().Get(email).Return(account, nil)
		mockEncrypter.EXPECT().Compare(password, account.Password).Return(true)
		mockResetRefreshToken.EXPECT().Reset(account.Id).Return(refreshToken, nil)

		httpRequest := createHttpRequest(t, email, password)
		httpResponse := signInController.Handle(*httpRequest)

		require.Equal(t, httpResponse.StatusCode, http.StatusOK)
		var responseBody controllers.SignInControllerResponse
		err := json.NewDecoder(httpResponse.Body).Decode(&responseBody)
		require.NoError(t, err)

		correctSignInControllerResponse := &controllers.SignInControllerResponse{
			RefreshToken: refreshToken,
		}
		require.Equal(t, &responseBody, correctSignInControllerResponse)
	})

	t.Run("InvalidBodyRequest", func(t *testing.T) {
		signInController, _, _, _, ctrl := setupMocks(t)

		defer ctrl.Finish()

		httpRequest := &protocols.HttpRequest{
			Body:   io.NopCloser(strings.NewReader("{invalid json")),
			Header: nil,
		}

		httpResponse := signInController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid body request")
	})

	t.Run("InvalidValidationEmailError", func(t *testing.T) {
		signInController, _, _, _, ctrl := setupMocks(t)
		defer ctrl.Finish()

		httpRequest := createHttpRequest(t, "invalid_email", password)
		httpResponse := signInController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusUnprocessableEntity, "Key: 'SignInControllerBody.Email' Error:Field validation for 'Email' failed on the 'email' tag")
	})

	t.Run("InvalidValidationPasswordError", func(t *testing.T) {
		signInController, _, _, _, ctrl := setupMocks(t)
		defer ctrl.Finish()

		httpMinPasswordRequest := createHttpRequest(t, email, "invalid")
		httpMaxPasswordRequest := createHttpRequest(t, email, strings.Repeat("invalid_password", 80))
		httpResponseMin := signInController.Handle(*httpMinPasswordRequest)
		httpResponseMax := signInController.Handle(*httpMaxPasswordRequest)

		verifyHttpResponse(t, httpResponseMin, http.StatusUnprocessableEntity, "Key: 'SignInControllerBody.Password' Error:Field validation for 'Password' failed on the 'min' tag")
		verifyHttpResponse(t, httpResponseMax, http.StatusUnprocessableEntity, "Key: 'SignInControllerBody.Password' Error:Field validation for 'Password' failed on the 'max' tag")
	})

	t.Run("InvalidEmailCredentials", func(t *testing.T) {
		signInController, _, mockGetAccountByEmail, _, ctrl := setupMocks(t)
		defer ctrl.Finish()

		mockGetAccountByEmail.EXPECT().Get(email).Return(nil, errors.New("fake-error"))

		httpRequest := createHttpRequest(t, email, password)
		httpResponse := signInController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid credentials")
	})

	t.Run("InvalidPasswordCredentials", func(t *testing.T) {
		signInController, mockEncrypter, mockGetAccountByEmail, _, ctrl := setupMocks(t)
		defer ctrl.Finish()

		account := &models.AccountModel{
			Id:           "fake-account-id",
			Email:        "fake-account-email@example.com",
			Password:     "fake-account-password",
			RefreshToken: refreshToken,
		}

		mockGetAccountByEmail.EXPECT().Get(email).Return(account, nil)
		mockEncrypter.EXPECT().Compare(password, account.Password).Return(false)

		httpRequest := createHttpRequest(t, email, password)
		httpResponse := signInController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid credentials")
	})

	t.Run("ResettingRefreshTokenError", func(t *testing.T) {
		signInController, mockEncrypter, mockGetAccountByEmail, mockResetRefreshToken, ctrl := setupMocks(t)
		defer ctrl.Finish()

		account := &models.AccountModel{
			Id:           "fake-account-id",
			Email:        "fake-account-email@example.com",
			Password:     "fake-account-password",
			RefreshToken: "fake-account-refresh-token",
		}

		mockGetAccountByEmail.EXPECT().Get(email).Return(account, nil)
		mockEncrypter.EXPECT().Compare(password, account.Password).Return(true)
		mockResetRefreshToken.EXPECT().Reset(account.Id).Return("", errors.New("fake-error"))

		httpRequest := createHttpRequest(t, email, password)
		httpResponse := signInController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "an error ocurred while resetting refresh token")
	})
}
