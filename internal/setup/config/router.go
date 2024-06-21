package config

import (
	"database/sql"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/routes"
)

func SetupRoutes(server *http.ServeMux, db *sql.DB) {
	routes.SignInRouter(server, db)
	routes.SignUpRouter(server, db)
	routes.RefreshTokenRouter(server, db)
}
