package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	email    = "test@example.com"
	password = "testpassword"
)

func setupMocks(t *testing.T) (*controllers.SignUpController, *mocks.MockGetAccountByEmail, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockGetAccountByEmail := mocks.NewMockGetAccountByEmail(ctrl)
	validate := validator.New(validator.WithRequiredStructEnabled())

	signUpController := &controllers.SignUpController{
		GetAccountByEmail: mockGetAccountByEmail,
		Validate:          validate,
	}

	return signUpController, mockGetAccountByEmail, ctrl
}

func createHttpRequest(t *testing.T, email, password string) *protocols.HttpRequest {
	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(&controllers.SignUpControllerBody{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)

	return &protocols.HttpRequest{
		Body:   io.NopCloser(&requestBody),
		Header: nil,
	}
}

func verifyHttpResponse(t *testing.T, httpResponse *protocols.HttpResponse, expectedStatusCode int, expectedError string) {
	require.Equal(t, httpResponse.StatusCode, expectedStatusCode)

	var responseBody protocols.ErrorResponse
	err := json.NewDecoder(httpResponse.Body).Decode(&responseBody)
	require.NoError(t, err)
	require.Equal(t, responseBody.Error, expectedError)
}

func TestSignUpController(t *testing.T) {
	t.Run("InvalidValidationEmailError", func(t *testing.T) {
		signUpController, _, ctrl := setupMocks(t)
		defer ctrl.Finish()

		httpRequest := createHttpRequest(t, "invalid_email", password)
		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusUnprocessableEntity, "Key: 'SignUpControllerBody.Email' Error:Field validation for 'Email' failed on the 'email' tag")
	})

	t.Run("InvalidValidationPasswordError", func(t *testing.T) {
		signUpController, _, ctrl := setupMocks(t)
		defer ctrl.Finish()

		httpMinPasswordRequest := createHttpRequest(t, email, "invalid")
		httpMaxPasswordRequest := createHttpRequest(t, email, strings.Repeat("invalid_password", 80))
		httpResponseMin := signUpController.Handle(*httpMinPasswordRequest)
		httpResponseMax := signUpController.Handle(*httpMaxPasswordRequest)

		verifyHttpResponse(t, httpResponseMin, http.StatusUnprocessableEntity, "Key: 'SignUpControllerBody.Password' Error:Field validation for 'Password' failed on the 'min' tag")
		verifyHttpResponse(t, httpResponseMax, http.StatusUnprocessableEntity, "Key: 'SignUpControllerBody.Password' Error:Field validation for 'Password' failed on the 'max' tag")
	})

	t.Run("InvalidEmailCredentials", func(t *testing.T) {
		signUpController, mockGetAccountByEmail, ctrl := setupMocks(t)
		defer ctrl.Finish()

		mockGetAccountByEmail.EXPECT().Get(email).Return(nil, nil)

		httpRequest := createHttpRequest(t, email, password)
		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusConflict, "User already exists")
	})
}
