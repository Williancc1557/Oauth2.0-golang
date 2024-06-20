package usecase_test

import (
	"errors"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMocks(t *testing.T) (*usecase.DbGetAccountByEmail, *mocks.MockGetAccountByEmailRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockGetAccountByEmailRepository := mocks.NewMockGetAccountByEmailRepository(ctrl)

	dbGetAccountByEmail := usecase.NewGetAccountByEmail(mockGetAccountByEmailRepository)

	return dbGetAccountByEmail, mockGetAccountByEmailRepository, ctrl
}

func TestDbGetAccountByEmail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		dbGetAccountByEmail, mockGetAccountByEmailRepository, ctrl := setupMocks(t)
		defer ctrl.Finish()

		email := "fake-email@example.com"

		account := &models.AccountModel{
			Id:           "fake-account-id",
			Email:        email,
			Password:     "fake-account-password",
			RefreshToken: "fake-account-refresh-token",
		}

		mockGetAccountByEmailRepository.EXPECT().Get(email).Return(account, nil)

		dbResponse, err := dbGetAccountByEmail.Get(email)

		require.NoErrorf(t, err, "an error ocurred while getting account: %v", err)
		require.NotNil(t, dbResponse)
		require.Equal(t, dbResponse, account)
	})

	t.Run("EmailRepositoryError", func(t *testing.T) {
		dbGetAccountByEmail, mockGetAccountByEmailRepository, ctrl := setupMocks(t)
		defer ctrl.Finish()

		email := "fake-email@example.com"

		mockGetAccountByEmailRepository.EXPECT().Get(email).Return(nil, errors.New("an error"))

		dbResponse, err := dbGetAccountByEmail.Get(email)

		require.Errorf(t, err, "got %v, want %v", err, errors.New("an error"))
		require.Nil(t, dbResponse, "got %v, want %v", dbResponse, nil)
	})
}
