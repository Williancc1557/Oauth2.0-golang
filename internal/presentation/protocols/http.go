package protocols

type HttpRequest struct {
	Body   any
	Header any
}

type HttpResponse struct {
	Body       any
	StatusCode any
}
