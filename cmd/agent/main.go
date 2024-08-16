package main

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"metrics/internal/configs"
	"metrics/internal/utils"
	"runtime"
	"time"
)

func main() {
	var cfg configs.AgentConfig

	flag.StringVar(&cfg.ServerAddress, "a", "http://localhost:8080", "address of the HTTP server endpoint (default localhost:8080)")
	flag.Int64Var(&cfg.PollInterval, "p", 2, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Int64Var(&cfg.ReportInterval, "r", 10, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Parse()

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	configs.ValidateAgentConfig(&cfg)

	gaugeMetrics := make(map[string]float64)
	counterMetrics := make(map[string]int64)

	collectInterval := time.Duration(cfg.PollInterval) * time.Second
	sendInterval := time.Duration(cfg.ReportInterval) * time.Second

	go func() {
		for {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			utils.CollectData(&memStats, gaugeMetrics, counterMetrics)
			time.Sleep(collectInterval)
		}
	}()

	for {
		utils.SendData(gaugeMetrics, counterMetrics, cfg.ServerAddress)
		time.Sleep(sendInterval)
	}
}
