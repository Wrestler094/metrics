package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"metrics/internal/handlers"
	"net/http"
)

func InitRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.GetMetricsHandler)
	router.Get("/value/{type}/{name}", handlers.GetMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", handlers.UpdateMetricHandler)

	return router
}

func main() {
	log.Fatalln(http.ListenAndServe(":8080", InitRouter()))
}
