package main

import (
	"flag"
	"go.uber.org/zap"
	"log"
	"metrics/internal/handlers"
	"metrics/internal/logger"
	"metrics/internal/storage"
	"net/http"
	"os"
)

func main() {
	var (
		flagRunAddress string
		flagLogLevel   string
	)

	// PARSE FLAGS
	flag.StringVar(&flagRunAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&flagLogLevel, "l", "info", "log level")
	flag.Parse()

	if envRunAddress := os.Getenv("ADDRESS"); envRunAddress != "" {
		flagRunAddress = envRunAddress
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagLogLevel = envLogLevel
	}

	// INITIALISE LOGGER
	err := logger.Initialize(flagLogLevel)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Log.Sync()

	// INITIALISE REPO & ROUTER
	var repo storage.Repository = storage.NewMemStorage()
	baseHandler := handlers.NewBaseHandler(repo)
	router := baseHandler.Router()

	// RUNNING SERVER
	logger.Log.Info("Running server", zap.String("address", flagRunAddress))

	if err := http.ListenAndServe(flagRunAddress, router); err != nil {
		logger.Log.Info("Server error", zap.Error(err))
	}
}
