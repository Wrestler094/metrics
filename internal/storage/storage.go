package storage

type MetricTypes struct {
	gauge   map[string]float64
	counter map[string]int64
}

type MemStorage struct {
	metrics MetricTypes
}

var Storage = MemStorage{
	metrics: MetricTypes{
		gauge:   map[string]float64{},
		counter: map[string]int64{},
	},
}

func (ms MemStorage) ChangeGauge(gaugeName string, newValue float64) {
	ms.metrics.gauge[gaugeName] = newValue
}

func (ms MemStorage) IncreaseCounter(counterName string, value int64) {
	res, ok := ms.metrics.counter[counterName]

	if !ok {
		ms.metrics.counter[counterName] = value
	} else {
		ms.metrics.counter[counterName] = res + value
	}
}
