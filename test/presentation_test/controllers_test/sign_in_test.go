package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"

	"github.com/golang/mock/gomock"
)

const (
	email        = "test@example.com"
	password     = "testpassword"
	refreshToken = "fake-refresh-token"
)

func setupMocks(t *testing.T) (*controllers.SignInController, *mocks.MockEncrypter, *mocks.MockGetAccountByEmail, *mocks.MockResetRefreshToken, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockEncrypter := mocks.NewMockEncrypter(ctrl)
	mockGetAccountByEmail := mocks.NewMockGetAccountByEmail(ctrl)
	mockResetRefreshToken := mocks.NewMockResetRefreshToken(ctrl)

	signInController := &controllers.SignInController{
		GetAccountByEmail: mockGetAccountByEmail,
		Encrypter:         mockEncrypter,
		ResetRefreshToken: mockResetRefreshToken,
	}

	return signInController, mockEncrypter, mockGetAccountByEmail, mockResetRefreshToken, ctrl
}

func createHttpRequest(t *testing.T, email, password string) *protocols.HttpRequest {
	var requestBody bytes.Buffer
	err := json.NewEncoder(&requestBody).Encode(&controllers.SignInControllerBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		t.Fatalf("an error occurred while encoding body: %v", err)
	}

	return &protocols.HttpRequest{
		Body:   io.NopCloser(&requestBody),
		Header: nil,
	}
}

func verifyHttpResponse(t *testing.T, httpResponse *protocols.HttpResponse, expectedStatusCode int, expectedError string) {
	if httpResponse.StatusCode != expectedStatusCode {
		t.Errorf("unexpected status code: got %d want %d", httpResponse.StatusCode, expectedStatusCode)
	}

	var responseBody protocols.ErrorResponse
	err := json.NewDecoder(httpResponse.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("an error occurred while decoding response body: %v", err)
	}

	if responseBody.Error != expectedError {
		t.Fatalf("unexpected error: got '%v' want '%v'", responseBody.Error, expectedError)
	}
}

func TestSignInController(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		signInController, mockEncrypter, mockGetAccountByEmail, mockResetRefreshToken, ctrl := setupMocks(t)
		defer ctrl.Finish()

		account := &models.AccountModel{
			Id:           "fake-account-id",
			Name:         "fake-account-name",
			Email:        "fake-account-email",
			Password:     "fake-account-password",
			RefreshToken: refreshToken,
		}

		mockGetAccountByEmail.EXPECT().Get(email).Return(account, nil)
		mockEncrypter.EXPECT().Compare(password, account.Password).Return(true)
		mockResetRefreshToken.EXPECT().Reset(account.Id).Return(refreshToken, nil)

		httpRequest := createHttpRequest(t, email, password)
		httpResponse := signInController.Handle(*httpRequest)

		if httpResponse.StatusCode != http.StatusOK {
			t.Errorf("unexpected status code: got %v want %v", httpResponse.StatusCode, http.StatusOK)
		}

		var responseBody controllers.SignInControllerResponse
		err := json.NewDecoder(httpResponse.Body).Decode(&responseBody)
		if err != nil {
			t.Fatalf("an error occurred while decoding response body: %v", err)
		}

		correctSignInControllerResponse := &controllers.SignInControllerResponse{
			RefreshToken: refreshToken,
		}
		if !reflect.DeepEqual(&responseBody, correctSignInControllerResponse) {
			t.Errorf("unexpected refresh token: got %v want %v", responseBody.RefreshToken, refreshToken)
		}
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
			Name:         "fake-account-name",
			Email:        "fake-account-email",
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
			Name:         "fake-account-name",
			Email:        "fake-account-email",
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
