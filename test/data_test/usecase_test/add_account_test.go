package usecase_test

import (
	"errors"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	usecaseDomain "github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupAddAccountMocks(t *testing.T) (*usecase.DbAddAccount, *mocks.MockAddAccountRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockAddAccountRepository := mocks.NewMockAddAccountRepository(ctrl)

	dbAddAccount := usecase.NewDbAddAccount(mockAddAccountRepository)

	return dbAddAccount, mockAddAccountRepository, ctrl
}

const (
	email    = "fake-email@example.com"
	password = "fake-password"
)

func TestAddAccountTest(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dbAddAccount, mockAddAccountRepository, ctrl := setupAddAccountMocks(t)
		defer ctrl.Finish()

		account := &protocols.AddAccountRepositoryOutput{
			Id:           "fake-account-id",
			Email:        email,
			Password:     "fake-account-password",
			RefreshToken: "fake-account-refresh-token",
		}

		mockAddAccountRepository.EXPECT().Add(&protocols.AddAccountRepositoryInput{
			Email:    email,
			Password: password,
		}).Return(account, nil)

		accountInput := &usecaseDomain.AddAccountInput{
			Email:    email,
			Password: password,
		}

		response, err := dbAddAccount.Add(accountInput)

		require.NoErrorf(t, err, "an error occurred while adding account: %v", err)
		require.NotNil(t, response)
		require.Equal(t, response.Id, account.Id)
		require.Equal(t, response.Email, account.Email)
		require.Equal(t, response.Password, account.Password)
		require.Equal(t, response.RefreshToken, account.RefreshToken)
	})

	t.Run("AddAccountRepositoryError", func(t *testing.T) {
		dbAddAccount, mockAddAccountRepository, ctrl := setupAddAccountMocks(t)
		defer ctrl.Finish()

		mockAddAccountRepository.EXPECT().Add(&protocols.AddAccountRepositoryInput{
			Email:    email,
			Password: password,
		}).Return(nil, errors.New("fake-error"))

		accountInput := &usecaseDomain.AddAccountInput{
			Email:    email,
			Password: password,
		}

		response, err := dbAddAccount.Add(accountInput)

		require.Error(t, err)
		require.Nil(t, response)
	})
}
