package main

import (
	"example/internal/setup"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", setup.Server())
}
