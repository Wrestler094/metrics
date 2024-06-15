package main

import (
	"fmt"
	"metrics/internal/utils"
	"runtime"
	"time"
)

const pollInterval = 2
const reportInterval = 10
const server = "http://localhost:8080"

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
	var memStats runtime.MemStats
	var tickNumber = 0
	var sendInterval = reportInterval / pollInterval

	for {
		runtime.ReadMemStats(&memStats)
		utils.CollectData(&memStats, gaugeMetrics, counterMetrics)

		if tickNumber != sendInterval {
			tickNumber++
		} else {
			utils.SendData(gaugeMetrics, counterMetrics, server)
			tickNumber = 1
		}
		fmt.Println(gaugeMetrics)
		time.Sleep(pollInterval * time.Second)
	}
}
