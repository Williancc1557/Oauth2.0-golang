package routes

import (
	"fmt"
	"net/http"
)

func SignInRouter(server *http.ServeMux) {
	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the return")
	})
}
