package controller_test

import (
	"bytes"
	"encoding/json"
	"example/internal/domain/models"
	"example/internal/presentation/controllers"
	"example/internal/presentation/protocols"
	"io"
	"testing"
)

type MockGetAccountByEmail struct{}

func (m *MockGetAccountByEmail) Get(email string) (*models.AccountModel, error) {
	return &models.AccountModel{
		Id:           "fake-account-id",
		Name:         "fake-account-name",
		Email:        "fake-account-email",
		Password:     "fake-account-password",
		RefreshToken: "fake-account-refresh-token",
	}, nil
}

type MockEncrypter struct{}

func (m *MockEncrypter) Hash(value string) (string, error) {
	return "fake-hash", nil
}

func (m *MockEncrypter) Compare(value string, hashedValue string) bool {
	return true
}

type MockResetRefreshToken struct{}

func (m *MockResetRefreshToken) Reset(userId string) (string, error) {
	return "fake-refresh-token", nil
}

func TestSignInController(t *testing.T) {
	mockGetAccountByEmail := &MockGetAccountByEmail{}
	mockEncrypter := &MockEncrypter{}
	MockResetRefreshToken := &MockResetRefreshToken{}

	signInController := controllers.SignInController{
		GetAccountByEmail: mockGetAccountByEmail,
		Encrypter:         mockEncrypter,
		ResetRefreshToken: MockResetRefreshToken,
	}

	value, err := json.Marshal(&controllers.SignInControllerBody{
		Email:    "WILL",
		Password: "aaaaa",
	})

	if err != nil {
		t.Errorf("an error ocurred while marshaling body: %v", err)
	}

	httpRequest := &protocols.HttpRequest{
		Body:   io.NopCloser(bytes.NewReader(value)),
		Header: nil,
	}

	signInController.Handle(*httpRequest)
}
