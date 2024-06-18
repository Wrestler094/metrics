package main

import (
	"flag"
	"fmt"
	"metrics/internal/handlers"
	storage "metrics/internal/storage"
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

	rep := storage.NewMemStorage()
	baseHandler := handlers.NewBaseHandler(rep)
	router := baseHandler.Router()

	fmt.Printf("Env Address: %s\n", os.Getenv("ADDRESS"))
	fmt.Printf("Starting server on %s\n", flagRunAddress)

	if err := http.ListenAndServe(flagRunAddress, router); err != nil {
		fmt.Printf("Server error - %s", err)
	}
}
