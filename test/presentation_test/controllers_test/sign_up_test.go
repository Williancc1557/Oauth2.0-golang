package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const ()

func setupSignUpMocks(t *testing.T) (*controllers.SignUpController, *mocks.MockGetAccountByEmail, *mocks.MockAddAccount, *mocks.MockCreateAccessToken, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockGetAccountByEmail := mocks.NewMockGetAccountByEmail(ctrl)
	mockAddAccount := mocks.NewMockAddAccount(ctrl)
	createAccessToken := mocks.NewMockCreateAccessToken(ctrl)

	signUpController := controllers.NewSignUpController(
		mockGetAccountByEmail,
		mockAddAccount,
		createAccessToken,
	)

	return signUpController, mockGetAccountByEmail, mockAddAccount, createAccessToken, ctrl
}

func createSignUpHttpRequest(t *testing.T, email, password string) *protocols.HttpRequest {
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

func convertReadCloserToStruct(reader io.ReadCloser, v interface{}) error {
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func TestSignUpController(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		signUpController, mockGetAccountByEmail, mockAddAccount, createAccessToken, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		signUpControllerInput := &usecase.AddAccountInput{
			Email:    email,
			Password: password,
		}
		AddAccountOutput := &usecase.AddAccountOutput{
			Id:           userId,
			Email:        email,
			Password:     password,
			RefreshToken: "fake-refresh-token",
		}
		CreateAccessTokenOutput := &usecase.CreateAccessTokenOutput{
			AccessToken: "fake-access-token",
			ExpiresIn:   123,
		}
		mockGetAccountByEmail.EXPECT().Get(email).Return(nil, errors.New("fake-error"))
		mockAddAccount.EXPECT().Add(signUpControllerInput).Return(AddAccountOutput, nil)
		createAccessToken.EXPECT().Create(userId).Return(CreateAccessTokenOutput, nil)

		httpRequest := createSignUpHttpRequest(t, email, password)
		httpResponse := signUpController.Handle(*httpRequest)

		require.Equal(t, httpResponse.StatusCode, http.StatusOK)

		expectedBody := &controllers.SignUpControllerResponse{
			ExpiresIn:    123,
			AccessToken:  "fake-access-token",
			RefreshToken: "fake-refresh-token",
		}

		var httpResponseBody controllers.SignUpControllerResponse
		err := convertReadCloserToStruct(httpResponse.Body, &httpResponseBody)
		require.NoError(t, err)
		require.Equal(t, expectedBody, &httpResponseBody)
	})

	t.Run("InvalidBodyRequest", func(t *testing.T) {
		signUpController, _, _, _, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		httpRequest := &protocols.HttpRequest{
			Body:   io.NopCloser(strings.NewReader("{invalid json")),
			Header: nil,
		}

		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid body request")
	})

	t.Run("InvalidValidationEmailError", func(t *testing.T) {
		signUpController, _, _, _, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		httpRequest := createSignUpHttpRequest(t, "invalid_email", password)
		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusUnprocessableEntity, "Key: 'SignUpControllerBody.Email' Error:Field validation for 'Email' failed on the 'email' tag")
	})

	t.Run("InvalidValidationPasswordError", func(t *testing.T) {
		signUpController, _, _, _, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		httpMinPasswordRequest := createSignUpHttpRequest(t, email, "invalid")
		httpMaxPasswordRequest := createSignUpHttpRequest(t, email, strings.Repeat("invalid_password", 80))
		httpResponseMin := signUpController.Handle(*httpMinPasswordRequest)
		httpResponseMax := signUpController.Handle(*httpMaxPasswordRequest)

		verifyHttpResponse(t, httpResponseMin, http.StatusUnprocessableEntity, "Key: 'SignUpControllerBody.Password' Error:Field validation for 'Password' failed on the 'min' tag")
		verifyHttpResponse(t, httpResponseMax, http.StatusUnprocessableEntity, "Key: 'SignUpControllerBody.Password' Error:Field validation for 'Password' failed on the 'max' tag")
	})

	t.Run("InvalidEmailCredentials", func(t *testing.T) {
		signUpController, mockGetAccountByEmail, _, _, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		mockGetAccountByEmail.EXPECT().Get(email).Return(nil, nil)

		httpRequest := createSignUpHttpRequest(t, email, password)
		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusConflict, "User already exists")
	})

	t.Run("ErrorWhileAddAccount", func(t *testing.T) {
		signUpController, mockGetAccountByEmail, mockAddAccount, _, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		signUpControllerInput := &usecase.AddAccountInput{
			Email:    email,
			Password: password,
		}
		mockGetAccountByEmail.EXPECT().Get(email).Return(nil, errors.New("fake-error"))
		mockAddAccount.EXPECT().Add(signUpControllerInput).Return(nil, errors.New("fake-error"))

		httpRequest := createSignUpHttpRequest(t, email, password)
		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "An error ocurred while adding account")
	})

	t.Run("ErrorWhileCreateAccess", func(t *testing.T) {
		signUpController, mockGetAccountByEmail, mockAddAccount, createAccessToken, ctrl := setupSignUpMocks(t)
		defer ctrl.Finish()

		signUpControllerInput := &usecase.AddAccountInput{
			Email:    email,
			Password: password,
		}
		AddAccountOutput := &usecase.AddAccountOutput{
			Id:           userId,
			Email:        email,
			Password:     password,
			RefreshToken: "fake-refresh-token",
		}
		mockGetAccountByEmail.EXPECT().Get(email).Return(nil, errors.New("fake-error"))
		mockAddAccount.EXPECT().Add(signUpControllerInput).Return(AddAccountOutput, nil)
		createAccessToken.EXPECT().Create(userId).Return(nil, errors.New("fake-error"))

		httpRequest := createSignUpHttpRequest(t, email, password)
		httpResponse := signUpController.Handle(*httpRequest)

		verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "An error ocurred while creating access token")
	})
}
