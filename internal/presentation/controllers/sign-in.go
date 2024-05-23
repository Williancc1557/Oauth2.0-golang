package controllers

import (
	"encoding/json"
	"example/internal/presentation/protocols"
	"net/http"
)

type SignInController struct{}

type a struct {
	Name string `json:"name"`
}

func (c *SignInController) Handle(r protocols.HttpRequest) (*protocols.HttpResponse, error) {
	var test a

	err := json.NewDecoder(r.Body).Decode(&test)

	if err != nil {
		return nil, err
	}

	return &protocols.HttpResponse{
		Body:       test,
		StatusCode: http.StatusOK,
	}, nil
}
