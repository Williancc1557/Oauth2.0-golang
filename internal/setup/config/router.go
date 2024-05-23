package config

import (
	"example/internal/setup/routes"
	"net/http"
)

func SetupRoutes(server *http.ServeMux) {
	routes.SignInRouter(server)
}
