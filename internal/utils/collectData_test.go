package utils

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestCollectData(t *testing.T) {
	type args struct {
		memStats       *runtime.MemStats
		gaugeMetrics   map[string]float64
		counterMetrics map[string]int64
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Random value set",
			args: args{
				memStats:       &runtime.MemStats{},
				gaugeMetrics:   map[string]float64{},
				counterMetrics: map[string]int64{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			CollectData(test.args.memStats, test.args.gaugeMetrics, test.args.counterMetrics)
			assert.NotEqual(t, 0.0, test.args.gaugeMetrics["RandomValue"], "Random value should be set")
		})
	}
}
