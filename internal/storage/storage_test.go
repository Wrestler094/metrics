package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeGauge(t *testing.T) {
	type metric struct {
		metricName string
		value      float64
	}

	tests := []struct {
		name    string
		storage *MemStorage
		metric  metric
	}{
		{
			name: "Write gauge metric in empty storage",
			storage: &MemStorage{
				Gauge:   map[string]float64{},
				Counter: map[string]int64{},
			},
			metric: metric{
				metricName: "testGauge",
				value:      100,
			},
		}, {
			name: "Write gauge metric in existed metric",
			storage: &MemStorage{
				Gauge: map[string]float64{
					"testGauge": 200,
				},
				Counter: map[string]int64{},
			},
			metric: metric{
				metricName: "testGauge",
				value:      100,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.storage.SetGaugeMetric(test.metric.metricName, test.metric.value)

			assert.Contains(t, test.storage.Gauge, test.metric.metricName)
			assert.Equal(t, test.metric.value, test.storage.Gauge[test.metric.metricName])
		})
	}
}

func TestIncreaseCounter(t *testing.T) {
	type metric struct {
		metricName string
		value      int64
	}

	tests := []struct {
		name    string
		storage *MemStorage
		metric  metric
		want    int64
	}{
		{
			name: "Write counter metric in empty storage",
			storage: &MemStorage{
				Gauge:   map[string]float64{},
				Counter: map[string]int64{},
			},
			metric: metric{
				metricName: "testCounter",
				value:      100,
			},
			want: 100,
		}, {
			name: "Write counter metric in existed metric",
			storage: &MemStorage{
				Gauge: map[string]float64{},
				Counter: map[string]int64{
					"testCounter": 200,
				},
			},
			metric: metric{
				metricName: "testCounter",
				value:      100,
			},
			want: 300,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.storage.SetCounterMetric(test.metric.metricName, test.metric.value)

			assert.Contains(t, test.storage.Counter, test.metric.metricName)
			assert.Equal(t, test.want, test.storage.Counter[test.metric.metricName])
		})
	}
}
