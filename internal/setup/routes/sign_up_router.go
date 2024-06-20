package routes

import (
	"database/sql"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/adapters"
	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/factory"
)

func SignUpRouter(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("POST /api/auth/sign-up", adapters.AdaptRoute(factory.MakeSignUpController(db)))
}
