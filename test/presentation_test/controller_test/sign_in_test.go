package controller_test

import (
	"bytes"
	"encoding/json"
	"example/internal/domain/models"
	"example/internal/presentation/controllers"
	"example/internal/presentation/protocols"
	"example/test/mocks"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSignInControllerSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEncrypter := mocks.NewMockEncrypter(ctrl)
	mockGetAccountByEmail := mocks.NewMockGetAccountByEmail(ctrl)
	mockResetRefreshToken := mocks.NewMockResetRefreshToken(ctrl)

	signInController := &controllers.SignInController{
		GetAccountByEmail: mockGetAccountByEmail,
		Encrypter:         mockEncrypter,
		ResetRefreshToken: mockResetRefreshToken,
	}

	email := "test@example.com"
	password := "testpassword"
	account := &models.AccountModel{
		Id:           "fake-account-id",
		Name:         "fake-account-name",
		Email:        "fake-account-email",
		Password:     "fake-account-password",
		RefreshToken: "fake-account-refresh-token",
	}
	refreshToken := "fake-refresh-token"

	mockGetAccountByEmail.EXPECT().Get(email).Return(account, nil)
	mockEncrypter.EXPECT().Compare(password, account.Password).Return(true)
	mockResetRefreshToken.EXPECT().Reset(account.Id).Return(refreshToken, nil)

	requestBody, err := json.Marshal(&controllers.SignInControllerBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		t.Fatalf("an error occurred while marshaling body: %v", err)
	}

	httpRequest := &protocols.HttpRequest{
		Body:   io.NopCloser(bytes.NewReader(requestBody)),
		Header: nil,
	}

	httpResponse := signInController.Handle(*httpRequest)

	if httpResponse.StatusCode >= 400 && httpResponse.StatusCode <= 500 {
		t.Errorf("unexpected status code: got %v want %v", httpResponse.StatusCode, http.StatusAccepted)
	}

	var responseBody controllers.SignInControllerResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&responseBody)

	if err != nil {
		t.Fatalf("an error occurred while decoding response body: %v", err)
	}

	t.Log(responseBody)

	if responseBody.RefreshToken != refreshToken {
		t.Errorf("unexpected refresh token: got %v want %v", responseBody.RefreshToken, refreshToken)
	}
}
