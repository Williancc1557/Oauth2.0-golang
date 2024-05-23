package adapters

import (
	"encoding/json"
	"example/internal/presentation/protocols"
	"net/http"
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

		res, err := a.Controller.Handle(*httpRequest)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(res.StatusCode)
		json.NewEncoder(w).Encode(res.Body)
	}
}
