package utils

import (
	"bytes"
	"encoding/json"
	"example/internal/presentation/protocols"
	"io"
)

func CreateResponse(body any, statusCode int) *protocols.HttpResponse {
	var bodyBuffer bytes.Buffer
	err := json.NewEncoder(&bodyBuffer).Encode(body)

	if err != nil {
		return &protocols.HttpResponse{
			Body:       io.NopCloser(&bodyBuffer),
			StatusCode: 400,
		}
	}

	return &protocols.HttpResponse{
		Body:       io.NopCloser(&bodyBuffer),
		StatusCode: statusCode,
	}
}
