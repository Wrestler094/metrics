package main

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"metrics/internal/utils"
	"runtime"
	"strings"
	"time"
)

type Config struct {
	ServerAddress  string `env:"ADDRESS"`
	PollInterval   int64  `env:"REPORT_INTERVAL"`
	ReportInterval int64  `env:"POLL_INTERVAL"`
}

func validateConfig(cfg *Config) {
	if cfg.PollInterval < 1 {
		cfg.PollInterval = 2
	}

	if cfg.ReportInterval < 1 {
		cfg.ReportInterval = 10
	}

	if !(strings.HasPrefix(cfg.ServerAddress, "http://")) {
		cfg.ServerAddress = "http://" + cfg.ServerAddress
	}
}

func main() {
	var cfg Config

	flag.StringVar(&cfg.ServerAddress, "a", "http://localhost:8080", "address of the HTTP server endpoint (default localhost:8080)")
	flag.Int64Var(&cfg.PollInterval, "p", 2, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Int64Var(&cfg.ReportInterval, "r", 10, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Parse()

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	validateConfig(&cfg)

	gaugeMetrics := make(map[string]float64)
	counterMetrics := make(map[string]int64)

	go func(g map[string]float64) {
		for {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			utils.CollectData(&memStats, g)
			time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
		}

	}(gaugeMetrics)
	go func(g map[string]float64) {
		for {
			utils.SendData(g, counterMetrics, cfg.ServerAddress)
			time.Sleep(time.Duration(cfg.ReportInterval) * time.Second)
		}

	}(gaugeMetrics)
	select {}
}
