package storage

import (
	"encoding/json"
	"go.uber.org/zap"
	"metrics/internal/logger"
	"os"
	"sync"
)

type Repository interface {
	RestoreData()
	SaveData()
	GetMetrics() (map[string]float64, map[string]int64)
	GetGaugeMetric(gaugeName string) (float64, bool)
	SetGaugeMetric(gaugeName string, newValue float64)
	GetCounterMetric(metricName string) (int64, bool)
	SetCounterMetric(metricName string, value int64)
}

type MemStorage struct {
	Gauge           map[string]float64 `json:"gauge"`
	Counter         map[string]int64   `json:"counter"`
	fileStoragePath string
	syncMode        bool
	mutex           sync.RWMutex
}

func NewMemStorage(fileStoragePath string, syncMode bool) *MemStorage {
	return &MemStorage{
		Gauge:           make(map[string]float64),
		Counter:         make(map[string]int64),
		fileStoragePath: fileStoragePath,
		syncMode:        syncMode,
	}
}

func (ms *MemStorage) RestoreData() {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	data, err := os.ReadFile(ms.fileStoragePath)
	if err != nil {
		logger.Log.Info("Error reading file:", zap.Error(err))
		return
	}
	if err := json.Unmarshal(data, ms); err != nil {
		logger.Log.Info("Error unmarshalling file:", zap.Error(err))
		return
	}

	logger.Log.Info("Storage data restored", zap.Any("data", string(data)))
}

func (ms *MemStorage) SaveData() {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	data, err := json.MarshalIndent(ms, "", "   ")
	if err != nil {
		logger.Log.Info("Error marshalling file:", zap.Error(err))
		return
	}

	err = os.WriteFile(ms.fileStoragePath, data, 0666)
	if err != nil {
		logger.Log.Info("Error writing file:", zap.Error(err))
		return
	}

	logger.Log.Info("Storage data saved", zap.Any("data", string(data)))
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
	ms.Gauge[gaugeName] = newValue
	ms.mutex.Unlock()

	if ms.syncMode {
		ms.SaveData()
	}
}

func (ms *MemStorage) GetCounterMetric(metricName string) (int64, bool) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	res, ok := ms.Counter[metricName]
	return res, ok
}

func (ms *MemStorage) SetCounterMetric(metricName string, value int64) {
	ms.mutex.Lock()
	ms.Counter[metricName] += value
	ms.mutex.Unlock()

	if ms.syncMode {
		ms.SaveData()
	}
}
