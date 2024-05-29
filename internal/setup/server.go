package setup

import (
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/config"
)

func Server() *http.ServeMux {
	mux := http.NewServeMux()

	config.SetupRoutes(mux)

	return mux
}
