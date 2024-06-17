package usecase_test

import (
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMocksForGetAccountByEmail(t *testing.T) (*usecase.DbResetRefreshToken, *mocks.MockResetRefreshToken, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockResetRefreshToken := mocks.NewMockResetRefreshToken(ctrl)

	return &usecase.DbResetRefreshToken{ResetRefreshTokenRepository: mockResetRefreshToken}, mockResetRefreshToken, ctrl
}

func TestDbResetRefreshToken(t *testing.T) {
	t.Run("ResetRefreshTokenSuccess", func(t *testing.T) {
		resetRefreshToken, mockResetRefreshToken, ctrl := setupMocksForGetAccountByEmail(t)

		defer ctrl.Finish()

		userId := "fake-user-id"
		refreshToken := "fake-refresh-token"
		mockResetRefreshToken.EXPECT().Reset(userId).Return(refreshToken, nil)

		value, err := resetRefreshToken.Reset(userId)

		require.NoError(t, err)
		require.Equal(t, refreshToken, value)
	})
}
