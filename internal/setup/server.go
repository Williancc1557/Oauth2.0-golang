package setup

import (
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/helpers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/config"
)

func Server() *http.ServeMux {
	mux := http.NewServeMux()

	db := helpers.PostgreHelper()

	config.SetupRoutes(mux, db)

	return mux
}
