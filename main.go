package main

import (
	"fmt"
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	port := ":8080"

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Some error occured. Err: %s", err)
		return
	}


	fmt.Println("server is running with port", port)

	hash, _ := bcrypt.GenerateFromPassword([]byte("willian123"), bcrypt.DefaultCost)
	fmt.Print(string(hash))

	err = http.ListenAndServe(port, setup.Server())

	if err != nil {
		fmt.Println("error ocurred: ", err.Error())
	}
}
