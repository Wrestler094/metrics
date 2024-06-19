package main

import (
	"flag"
	"fmt"
	"metrics/internal/handlers"
	"metrics/internal/storage"
	"net/http"
	"os"
)

var flagRunAddress string

func main() {
	flag.StringVar(&flagRunAddress, "a", "localhost:8080", "address and port to run server")
	flag.Parse()

	if envRunAddress := os.Getenv("ADDRESS"); envRunAddress != "" {
		flagRunAddress = envRunAddress
	}

	var repo storage.Repository = storage.NewMemStorage()
	baseHandler := handlers.NewBaseHandler(repo)
	router := baseHandler.Router()

	fmt.Printf("Starting server on %s\n", flagRunAddress)

	if err := http.ListenAndServe(flagRunAddress, router); err != nil {
		fmt.Printf("Server error - %s", err)
	}
}
