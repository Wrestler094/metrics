package main

import (
	"flag"
	"go.uber.org/zap"
	"metrics/internal/logger"
	"os"
	"strconv"
)

type Flags struct {
	flagRunAddress      string
	flagLogLevel        string
	flagStoreInterval   int64
	flagFileStoragePath string
	flagRestore         bool
}

type stringFlags struct {
	flagRunAddress      string
	flagLogLevel        string
	flagStoreInterval   string
	flagFileStoragePath string
	flagRestore         string
}

func (f *Flags) Parse() {
	var flags stringFlags
	flag.StringVar(&flags.flagRunAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&flags.flagLogLevel, "l", "info", "log level")
	flag.StringVar(&flags.flagStoreInterval, "i", "300", "time interval in seconds after which current server readings are saved to disk")
	flag.StringVar(&flags.flagFileStoragePath, "f", "internal/storage/storage.json", "path to the file where current values are saved")
	flag.StringVar(&flags.flagRestore, "r", "true", "load or not previously saved values from the specified file when the server starts")
	flag.Parse()

	if envRunAddress := os.Getenv("ADDRESS"); envRunAddress != "" {
		flags.flagRunAddress = envRunAddress
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flags.flagLogLevel = envLogLevel
	}
	if envStoreInterval := os.Getenv("STORE_INTERVAL"); envStoreInterval != "" {
		flags.flagStoreInterval = envStoreInterval
	}
	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		flags.flagFileStoragePath = envFileStoragePath
	}
	if envRestore := os.Getenv("RESTORE"); envRestore != "" {
		flags.flagRestore = envRestore
	}

	storeInterval, err := strconv.ParseInt(flags.flagStoreInterval, 10, 64)
	if err != nil {
		logger.Log.Info("Convert flagStoreInterval to int unsuccessful", zap.Error(err))
		storeInterval = 300
	}

	restore, err := strconv.ParseBool(flags.flagRestore)
	if err != nil {
		logger.Log.Info("Convert flagRestore to bool unsuccessful", zap.Error(err))
		restore = true
	}

	f.flagRunAddress = flags.flagRunAddress
	f.flagLogLevel = flags.flagLogLevel
	f.flagStoreInterval = storeInterval
	f.flagFileStoragePath = flags.flagFileStoragePath
	f.flagRestore = restore
}
