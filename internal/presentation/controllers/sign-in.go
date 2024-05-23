package controllers

import "example/internal/presentation/protocols"

type SignInController struct{}

func (c *SignInController) Handle(r protocols.HttpRequest) (protocols.HttpResponse, error) {
	return protocols.HttpResponse{}, nil
}
