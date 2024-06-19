package storage

type Repository interface {
	GetMetrics() (*map[string]float64, *map[string]int64)
	GetGaugeMetric(gaugeName string) (float64, bool)
	SetGaugeMetric(gaugeName string, newValue float64)
	GetCounterMetric(metricName string) (int64, bool)
	SetCounterMetric(metricName string, value int64)
}

type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
}

func (ms *MemStorage) GetMetrics() (*map[string]float64, *map[string]int64) {
	return &ms.Gauge, &ms.Counter
}

func (ms *MemStorage) GetGaugeMetric(metricName string) (float64, bool) {
	res, ok := ms.Gauge[metricName]
	return res, ok
}

func (ms *MemStorage) SetGaugeMetric(gaugeName string, newValue float64) {
	ms.Gauge[gaugeName] = newValue
}

func (ms *MemStorage) GetCounterMetric(metricName string) (int64, bool) {
	res, ok := ms.Counter[metricName]
	return res, ok
}

func (ms *MemStorage) SetCounterMetric(metricName string, value int64) {
	res, ok := ms.Counter[metricName]

	if !ok {
		ms.Counter[metricName] = value
	} else {
		ms.Counter[metricName] = res + value
	}

	//if _, ok := ms.Counter[metricName]; ok {
	//	ms.Counter[metricName] += value
	//	return
	//}
	//ms.Counter[metricName] = value
}
