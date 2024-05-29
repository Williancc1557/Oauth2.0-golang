package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/protocols"
)

type AdaptRouteDependencies struct {
	Controller protocols.Controller
}

func (a *AdaptRouteDependencies) AdaptRoute() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpRequest := &protocols.HttpRequest{
			Body:   r.Body,
			Header: r.Header,
		}

		res := a.Controller.Handle(*httpRequest)

		w.WriteHeader(res.StatusCode)
		json.NewEncoder(w).Encode(res.Body)
	}
}
