package utils_test

import (
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
	"github.com/stretchr/testify/require"
)

func SetupMocks(t *testing.T) *utils.CreateAccessTokenUtil {
	createAccessTokenUtil := utils.NewCreateAccessTokenUtil("fake-private-key")

	return createAccessTokenUtil
}

func TestCreateAccessTokenUtil(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		createAccessTokenUtil := SetupMocks(t)

		token, err := createAccessTokenUtil.Create("fake-id")
		require.NoError(t, err)
		require.NotNil(t, token)
	})
}
