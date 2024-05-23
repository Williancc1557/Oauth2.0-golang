package config

import (
	"fmt"
	"net/http"
)

func SetupRoutes(server *http.ServeMux) {
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the return")
	})
}
