package utils_test

import (
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func setupMocks() *utils.EncrypterUtil {
	encripter := &utils.EncrypterUtil{}

	return encripter
}

func TestEncrypterUtils(t *testing.T) {
	t.Run("HashSuccess", func(t *testing.T) {
		encripter := setupMocks()

		password := "123"
		hash, err := encripter.Hash(password)
		require.NoError(t, err)

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		require.NoError(t, err)
	})

	t.Run("CompareSuccess", func(t *testing.T) {
		encripter := setupMocks()

		password := "123"
		hash, err := encripter.Hash(password)
		require.NoError(t, err)

		value := encripter.Compare(password, hash)

		t.Log("aaaaaaaaaaaaa", value)
		require.True(t, value)
	})
}
