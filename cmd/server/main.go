package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"metrics/internal/handlers"
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

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.GetMetricsHandler)
	router.Get("/value/{type}/{name}", handlers.GetMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", handlers.PostMetricHandler)

	panic(http.ListenAndServe(flagRunAddress, router))
}
