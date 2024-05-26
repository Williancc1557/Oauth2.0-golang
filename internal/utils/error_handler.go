package utils

import (
	"errors"
	"example/internal/presentation/protocols"
)

func HandleError(text string, statusCode int) *protocols.ErrorResponse {
	return &protocols.ErrorResponse{
		StatusCode: statusCode,
		Error:      errors.New(text),
	}
}
