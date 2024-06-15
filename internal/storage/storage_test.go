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
		storage MemStorage
		metric  metric
	}{
		{
			name: "Write gauge metric in empty storage",
			storage: MemStorage{
				metrics: MetricTypes{
					gauge:   map[string]float64{},
					counter: map[string]int64{},
				},
			},
			metric: metric{
				metricName: "testGauge",
				value:      100,
			},
		}, {
			name: "Write gauge metric in existed metric",
			storage: MemStorage{
				metrics: MetricTypes{
					gauge: map[string]float64{
						"testGauge": 200,
					},
					counter: map[string]int64{},
				},
			},
			metric: metric{
				metricName: "testGauge",
				value:      100,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.storage.ChangeGauge(test.metric.metricName, test.metric.value)

			assert.Contains(t, test.storage.metrics.gauge, test.metric.metricName)
			assert.Equal(t, test.metric.value, test.storage.metrics.gauge[test.metric.metricName])
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
		storage MemStorage
		metric  metric
		want    int64
	}{
		{
			name: "Write counter metric in empty storage",
			storage: MemStorage{
				metrics: MetricTypes{
					gauge:   map[string]float64{},
					counter: map[string]int64{},
				},
			},
			metric: metric{
				metricName: "testCounter",
				value:      100,
			},
			want: 100,
		}, {
			name: "Write counter metric in existed metric",
			storage: MemStorage{
				metrics: MetricTypes{
					gauge: map[string]float64{},
					counter: map[string]int64{
						"testCounter": 200,
					},
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
			test.storage.IncreaseCounter(test.metric.metricName, test.metric.value)

			assert.Contains(t, test.storage.metrics.counter, test.metric.metricName)
			assert.Equal(t, test.want, test.storage.metrics.counter[test.metric.metricName])
		})
	}
}
