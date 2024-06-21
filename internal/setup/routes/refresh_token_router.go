package routes

import (
	"database/sql"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/adapters"
	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/factory"
)

func RefreshTokenRouter(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("GET /api/auth/refresh-token", adapters.AdaptRoute(factory.MakeRefreshTokenController(db)))
}
