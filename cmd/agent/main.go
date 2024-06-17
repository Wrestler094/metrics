package main

import (
	"flag"
	"metrics/internal/utils"
	"runtime"
	"time"
)

type Config struct {
	ServerAddress  string `env:"ADDRESS"`
	PollInterval   int64  `env:"REPORT_INTERVAL"`
	ReportInterval int64  `env:"POLL_INTERVAL"`
}

var gaugeMetrics = map[string]float64{
	"Alloc":         0,
	"BuckHashSys":   0,
	"Frees":         0,
	"GCCPUFraction": 0,
	"GCSys":         0,
	"HeapAlloc":     0,
	"HeapIdle":      0,
	"HeapInuse":     0,
	"HeapObjects":   0,
	"HeapReleased":  0,
	"HeapSys":       0,
	"LastGC":        0,
	"Lookups":       0,
	"MCacheInuse":   0,
	"MCacheSys":     0,
	"MSpanInuse":    0,
	"MSpanSys":      0,
	"Mallocs":       0,
	"NextGC":        0,
	"NumForcedGC":   0,
	"NumGC":         0,
	"OtherSys":      0,
	"PauseTotalNs":  0,
	"StackInuse":    0,
	"StackSys":      0,
	"Sys":           0,
	"TotalAlloc":    0,
	"RandomValue":   0,
}

var counterMetrics = map[string]int64{
	"PollCount": 0,
}

func main() {
	var cfg Config

	flag.StringVar(&cfg.ServerAddress, "a", "http://localhost:8080", "address of the HTTP server endpoint (default localhost:8080)")
	flag.Int64Var(&cfg.PollInterval, "p", 2, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Int64Var(&cfg.ReportInterval, "r", 10, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Parse()

	//err := env.Parse(&cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	utils.ValidateFlags(&cfg.PollInterval, &cfg.ReportInterval, &cfg.ServerAddress)

	var memStats runtime.MemStats
	var sendInterval = cfg.ReportInterval / cfg.PollInterval

	for {
		runtime.ReadMemStats(&memStats)
		utils.CollectData(&memStats, gaugeMetrics)

		if counterMetrics["PollCount"] != sendInterval {
			counterMetrics["PollCount"]++
		} else {
			utils.SendData(gaugeMetrics, counterMetrics, cfg.ServerAddress)
			counterMetrics["PollCount"] = 1
		}

		time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
	}
}
