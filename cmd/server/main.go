package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"metrics/internal/handlers"
	"net/http"
)

type Config struct {
	Address string `env:"ADDRESS"`
}

var flagRunAddress string

func main() {
	flag.StringVar(&flagRunAddress, "a", "localhost:8080", "address and port to run server")
	flag.Parse()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	if cfg.Address != "" {
		flagRunAddress = cfg.Address
	}
	//if envRunAddress := os.Getenv("ADDRESS"); envRunAddress != "" {
	//	flagRunAddress = envRunAddress
	//}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.GetMetricsHandler)
	router.Get("/value/{type}/{name}", handlers.GetMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", handlers.UpdateMetricHandler)

	log.Fatal(http.ListenAndServe(flagRunAddress, router))
}
