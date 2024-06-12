package routes

import (
	"database/sql"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/adapters"
	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/factory"
)

func SignInRouter(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("POST /", adapters.AdaptRoute(factory.MakeSignInController(db)))
}
