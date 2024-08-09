package storage

import "sync"

type Repository interface {
	GetMetrics() (map[string]float64, map[string]int64)
	GetGaugeMetric(gaugeName string) (float64, bool)
	SetGaugeMetric(gaugeName string, newValue float64)
	GetCounterMetric(metricName string) (int64, bool)
	SetCounterMetric(metricName string, value int64)
}

type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
	mutex   sync.RWMutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
}

func (ms *MemStorage) GetMetrics() (map[string]float64, map[string]int64) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	gaugeCopy := make(map[string]float64)
	counterCopy := make(map[string]int64)

	for key, value := range ms.Gauge {
		gaugeCopy[key] = value
	}

	for key, value := range ms.Counter {
		counterCopy[key] = value
	}

	return gaugeCopy, counterCopy
}

func (ms *MemStorage) GetGaugeMetric(metricName string) (float64, bool) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	res, ok := ms.Gauge[metricName]
	return res, ok
}

func (ms *MemStorage) SetGaugeMetric(gaugeName string, newValue float64) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.Gauge[gaugeName] = newValue
}

func (ms *MemStorage) GetCounterMetric(metricName string) (int64, bool) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	res, ok := ms.Counter[metricName]
	return res, ok
}

func (ms *MemStorage) SetCounterMetric(metricName string, value int64) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.Counter[metricName] += value
}
