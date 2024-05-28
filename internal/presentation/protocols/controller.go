package protocols

type Controller interface {
	Handle(HttpRequest) *HttpResponse
}
