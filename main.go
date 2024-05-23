package main

import (
	"example/internal/setup"
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"

	fmt.Println("server is running with port", port)

	err := http.ListenAndServe(port, setup.Server())

	if err != nil {
		fmt.Println("error ocurred: ", err.Error())
	}

}
