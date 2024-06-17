package utils_test

import (
	"errors"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func setupMocks(t *testing.T) (*utils.EncrypterUtil, *mocks.MockCryptoUtil, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockBcrypt := mocks.NewMockCryptoUtil(ctrl)

	encripter := &utils.EncrypterUtil{
		Crypto: mockBcrypt,
	}

	return encripter, mockBcrypt, ctrl
}

func TestEncrypterUtils(t *testing.T) {
	t.Run("HashSuccess", func(t *testing.T) {
		encripter, mockBcrypt, ctrl := setupMocks(t)
		ctrl.Finish()

		password := "1234567"
		validHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		require.NoError(t, err)
		mockBcrypt.EXPECT().GenerateFromPassword([]byte(password), bcrypt.DefaultCost).Return(validHash, nil)
		hash, err := encripter.Hash(password)
		require.NoError(t, err)

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		require.NoError(t, err)
	})

	t.Run("HashError", func(t *testing.T) {
		encripter, mockBcrypt, ctrl := setupMocks(t)
		ctrl.Finish()

		password := ""

		mockBcrypt.EXPECT().GenerateFromPassword([]byte(password), bcrypt.DefaultCost).Return(nil, errors.New("fake-error"))
		hash, err := encripter.Hash(password)
		require.Error(t, err)

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		require.Error(t, err)
	})

	t.Run("CompareSuccess", func(t *testing.T) {
		encripter, mockBcrypt, ctrl := setupMocks(t)
		ctrl.Finish()

		password := "123"
		mockBcrypt.EXPECT().CompareHashAndPassword([]byte(password), []byte(password)).Return(nil)

		value := encripter.Compare(password, password)

		require.True(t, value)
	})
}
