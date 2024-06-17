package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"metrics/internal/handlers"
	"metrics/internal/utils"
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

	utils.ValidateServerAddress(&flagRunAddress)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.GetMetricsHandler)
	router.Get("/value/{type}/{name}", handlers.GetMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", handlers.UpdateMetricHandler)

	log.Fatalln(http.ListenAndServe(flagRunAddress, router))
}
