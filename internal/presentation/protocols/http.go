package protocols

import (
	"io"
	"net/http"
)

type HttpRequest struct {
	Body   io.ReadCloser
	Header http.Header
}

type HttpResponse struct {
	Body       io.ReadCloser
	StatusCode int
}
