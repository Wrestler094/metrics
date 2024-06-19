package main

import (
	"metrics/internal/config"
	"metrics/internal/utils"
	"runtime"
	"time"
)

func main() {
	var cfg config.Config
	config.ParseAgentConfig(&cfg)
	config.ValidateAgentConfig(&cfg)

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
