package usecase_test

import (
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupGetAccountByRefreshTokenMocks(t *testing.T) (*usecase.DbGetAccountByRefreshToken, *mocks.MockGetAccountByRefreshTokenRepository, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockGetAccountByRefreshTokenRepository := mocks.NewMockGetAccountByRefreshTokenRepository(ctrl)

	sut := usecase.NewDbGetAccountByRefreshToken(mockGetAccountByRefreshTokenRepository)

	return sut, mockGetAccountByRefreshTokenRepository, ctrl
}

func TestDbGetAccountByRefreshToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, getAccountByRefreshTokenRepository, ctrl := setupGetAccountByRefreshTokenMocks(t)
		ctrl.Finish()

		account := &models.AccountModel{
			Id:           userId,
			Email:        email,
			Password:     password,
			RefreshToken: refreshToken,
		}

		getAccountByRefreshTokenRepository.EXPECT().Get("fake-refresh-token").Return(account, nil)

		account, err := sut.GetAccountByRefreshTokenRepository.Get(refreshToken)

		require.NoError(t, err)
		require.NotNil(t, account)
	})
}
