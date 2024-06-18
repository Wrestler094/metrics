package storage

type Repository interface {
	GetMetrics() (map[string]float64, map[string]int64)
	GetGaugeMetric(gaugeName string) (float64, bool)
	SetGaugeMetric(gaugeName string, newValue float64)
	GetCounterMetric(metricName string) (int64, bool)
	SetCounterMetric(metricName string, value int64)
}

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

var Storage = MemStorage{
	gauge:   map[string]float64{},
	counter: map[string]int64{},
}

func (ms MemStorage) GetMetrics() (map[string]float64, map[string]int64) {
	return ms.gauge, ms.counter
}

func (ms MemStorage) GetGaugeMetric(metricName string) (float64, bool) {
	res, ok := ms.gauge[metricName]
	return res, ok
}

func (ms MemStorage) SetGaugeMetric(gaugeName string, newValue float64) {
	ms.gauge[gaugeName] = newValue
}

func (ms MemStorage) GetCounterMetric(metricName string) (int64, bool) {
	res, ok := ms.counter[metricName]
	return res, ok
}

func (ms MemStorage) SetCounterMetric(metricName string, value int64) {
	res, ok := ms.counter[metricName]

	if !ok {
		ms.counter[metricName] = value
	} else {
		ms.counter[metricName] = res + value
	}
}
