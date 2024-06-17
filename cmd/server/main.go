package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"metrics/internal/handlers"
	"net/http"
	"os"
)

var flagRunAddress string

func TimerTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(r.Host + "\n")
		fmt.Printf(r.URL.Path + " " + r.URL.Host + " " + r.Method + "\n")

		next.ServeHTTP(w, r)

		fmt.Printf(r.Host + "\n")
		fmt.Printf(r.URL.Path + " " + r.URL.Host + " " + r.Method + "\n")
	})
}

func main() {
	flag.StringVar(&flagRunAddress, "a", ":8080", "address and port to run server")
	flag.Parse()

	if envRunAddress := os.Getenv("ADDRESS"); envRunAddress != "" {
		flagRunAddress = envRunAddress
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(TimerTrace)

	router.Get("/", handlers.GetMetricsHandler)
	router.Get("/value/{type}/{name}", handlers.GetMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", handlers.UpdateMetricHandler)

	log.Fatalln(http.ListenAndServe(flagRunAddress, router))
}
