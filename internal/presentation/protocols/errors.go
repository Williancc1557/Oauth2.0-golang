package protocols

type ErrorResponse struct {
	StatusCode int
	Error      error
}
