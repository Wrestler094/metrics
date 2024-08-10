package utils

import (
	"math/rand"
	"runtime"
)

func CollectData(memStats *runtime.MemStats, gaugeMetrics map[string]float64, counterMetrics map[string]int64) {
	gaugeMetrics["Alloc"] = float64(memStats.Alloc)
	gaugeMetrics["BuckHashSys"] = float64(memStats.BuckHashSys)
	gaugeMetrics["Frees"] = float64(memStats.Frees)
	gaugeMetrics["GCCPUFraction"] = memStats.GCCPUFraction
	gaugeMetrics["GCSys"] = float64(memStats.GCSys)
	gaugeMetrics["HeapAlloc"] = float64(memStats.HeapAlloc)
	gaugeMetrics["HeapIdle"] = float64(memStats.HeapIdle)
	gaugeMetrics["HeapInuse"] = float64(memStats.HeapInuse)
	gaugeMetrics["HeapObjects"] = float64(memStats.HeapObjects)
	gaugeMetrics["HeapReleased"] = float64(memStats.HeapReleased)
	gaugeMetrics["HeapSys"] = float64(memStats.HeapSys)
	gaugeMetrics["LastGC"] = float64(memStats.LastGC)
	gaugeMetrics["Lookups"] = float64(memStats.Lookups)
	gaugeMetrics["MCacheInuse"] = float64(memStats.MCacheInuse)
	gaugeMetrics["MCacheSys"] = float64(memStats.MCacheSys)
	gaugeMetrics["MSpanInuse"] = float64(memStats.MSpanInuse)
	gaugeMetrics["MSpanSys"] = float64(memStats.MSpanSys)
	gaugeMetrics["Mallocs"] = float64(memStats.Mallocs)
	gaugeMetrics["NextGC"] = float64(memStats.NextGC)
	gaugeMetrics["NumForcedGC"] = float64(memStats.NumForcedGC)
	gaugeMetrics["NumGC"] = float64(memStats.NumGC)
	gaugeMetrics["OtherSys"] = float64(memStats.OtherSys)
	gaugeMetrics["PauseTotalNs"] = float64(memStats.PauseTotalNs)
	gaugeMetrics["StackInuse"] = float64(memStats.StackInuse)
	gaugeMetrics["StackSys"] = float64(memStats.StackSys)
	gaugeMetrics["Sys"] = float64(memStats.Sys)
	gaugeMetrics["TotalAlloc"] = float64(memStats.TotalAlloc)
	gaugeMetrics["RandomValue"] = rand.Float64()
	counterMetrics["PollCount"] += 1
}
