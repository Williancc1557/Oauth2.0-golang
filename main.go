package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Williancc1557/Oauth2.0-golang/internal/setup"
	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/config"
)

func main() {
	port := ":8080"
	config.LoadEnvFile(".env")

	log.Println("server is running with port", port)

	sm := http.Server{
		Addr:         port,
		Handler:      setup.Server(),
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := sm.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	sm.Shutdown(tc)
}
