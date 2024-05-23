package routes

import (
	"example/internal/presentation/controllers"
	"example/internal/setup/adapters"
	"net/http"
)

func SignInRouter(server *http.ServeMux) {
	signInController := &controllers.SignInController{}
	a := adapters.AdaptRouteDependencies{
		Controller: signInController,
	}
	server.HandleFunc("POST /", a.AdaptRoute())
}
