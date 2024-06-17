package main

import (
	"flag"
	"fmt"
	"metrics/internal/utils"
	"runtime"
	"time"
)

var (
	flagServerAddress  string
	flagPollInterval   int64
	flagReportInterval int64
)

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
	flag.StringVar(&flagServerAddress, "a", "http://127.0.0.1:8080", "address of the HTTP server endpoint (default localhost:8080)")
	flag.Int64Var(&flagPollInterval, "p", 2, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Int64Var(&flagReportInterval, "r", 10, "frequency of sending metrics to the server (default 10 seconds)")
	flag.Parse()
	fmt.Println(flagServerAddress)

	if flagPollInterval < 1 {
		flagPollInterval = 2
	}

	if flagReportInterval < 1 {
		flagReportInterval = 10
	}

	var memStats runtime.MemStats
	var sendInterval = flagReportInterval / flagPollInterval

	for {
		runtime.ReadMemStats(&memStats)
		utils.CollectData(&memStats, gaugeMetrics)

		if counterMetrics["PollCount"] != sendInterval {
			counterMetrics["PollCount"]++
		} else {
			utils.SendData(gaugeMetrics, counterMetrics, flagServerAddress)
			counterMetrics["PollCount"] = 1
		}

		time.Sleep(time.Duration(flagPollInterval) * time.Second)
	}
}
