package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
	"github.com/stretchr/testify/require"
)

const (
	userId       = "fake-id"
	email        = "test@example.com"
	password     = "testpassword"
	refreshToken = "fake-refresh-token"
)

func verifyHttpResponse(t *testing.T, httpResponse *protocols.HttpResponse, expectedStatusCode int, expectedError string) {
	require.Equal(t, httpResponse.StatusCode, expectedStatusCode)

	var responseBody protocols.ErrorResponse
	err := json.NewDecoder(httpResponse.Body).Decode(&responseBody)
	require.NoError(t, err)
	require.Equal(t, responseBody.Error, expectedError)
}
