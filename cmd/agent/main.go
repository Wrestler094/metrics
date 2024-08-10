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

	go func(g map[string]float64, c map[string]int64) {
		for {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			utils.CollectData(&memStats, g, c)
			time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
		}

	}(gaugeMetrics, counterMetrics)

	for {
		utils.SendData(gaugeMetrics, counterMetrics, cfg.ServerAddress)
		time.Sleep(time.Duration(cfg.ReportInterval) * time.Second)
	}
}
