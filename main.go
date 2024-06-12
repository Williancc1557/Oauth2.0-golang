package main

import (
	"fmt"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup"
	"github.com/joho/godotenv"
)

func main() {
	port := ":8080"

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Some error occured. Err: %s", err)
		return
	}

	fmt.Println("server is running with port", port)

	err = http.ListenAndServe(port, setup.Server())

	if err != nil {
		fmt.Println("error ocurred: ", err.Error())
	}
}
