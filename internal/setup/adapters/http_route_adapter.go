package adapters

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
)


func AdaptRoute(controller protocols.Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpRequest := &protocols.HttpRequest{
			Body:   r.Body,
			Header: r.Header,
		}

		res := controller.Handle(*httpRequest)

		w.WriteHeader(res.StatusCode)
		_, err := io.Copy(w, res.Body)
		if err != nil {
			http.Error(w, "Failed to write response body", http.StatusInternalServerError)
			return
		}

		bodyBytes, _ := io.ReadAll(res.Body)
		fmt.Println(string(bodyBytes))
	}
}
