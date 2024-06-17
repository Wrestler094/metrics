package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"metrics/internal/handlers"
	"net/http"
)

var flagRunAddress string

func main() {
	flag.StringVar(&flagRunAddress, "a", ":8080", "address and port to run server")
	flag.Parse()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.GetMetricsHandler)
	router.Get("/value/{type}/{name}", handlers.GetMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", handlers.UpdateMetricHandler)

	log.Fatalln(http.ListenAndServe(flagRunAddress, router))
}
