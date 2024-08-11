package main

import (
	"go.uber.org/zap"
	"log"
	"metrics/internal/handlers"
	"metrics/internal/logger"
	"metrics/internal/storage"
	"net/http"
	"time"
)

func main() {
	// PARSE FLAGS
	flags := Flags{}
	flags.Parse()

	// INITIALISE LOGGER
	err := logger.Initialize(flags.flagLogLevel)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Log.Sync()

	// INITIALISE REPO & ROUTER
	var syncMode = flags.flagStoreInterval == 0
	var repo storage.Repository = storage.NewMemStorage(flags.flagFileStoragePath, syncMode)
	if flags.flagRestore {
		repo.RestoreData()
	}
	baseHandler := handlers.NewBaseHandler(repo)
	router := baseHandler.Router()

	// INIT SAVING GOROUTINE
	go func() {
		for !syncMode {
			time.Sleep(time.Duration(flags.flagStoreInterval) * time.Second)
			repo.SaveData()
		}
	}()

	// RUNNING SERVER
	logger.Log.Info("Running server", zap.String("address", flags.flagRunAddress))

	if err := http.ListenAndServe(flags.flagRunAddress, router); err != nil {
		logger.Log.Info("Server error", zap.Error(err))
	}
}
