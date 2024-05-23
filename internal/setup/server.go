package setup

import (
	"example/internal/config"
	"net/http"
)

func Server() *http.ServeMux {
	mux := http.NewServeMux()

	config.SetupRouters(mux)

	return mux
}
