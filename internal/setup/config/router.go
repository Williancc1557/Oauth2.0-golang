package config

import (
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/routes"
)

func SetupRoutes(server *http.ServeMux) {
	routes.SignInRouter(server)
}
