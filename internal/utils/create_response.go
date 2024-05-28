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
		var errorBuffer bytes.Buffer
		json.NewEncoder(&errorBuffer).Encode(err)

		return &protocols.HttpResponse{
			Body:       io.NopCloser(&errorBuffer),
			StatusCode: 400,
		}
	}

	return &protocols.HttpResponse{
		Body:       io.NopCloser(&bodyBuffer),
		StatusCode: statusCode,
	}
}
