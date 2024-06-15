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
			name: "Test 1",
			args: args{
				memStats:       &runtime.MemStats{},
				gaugeMetrics:   map[string]float64{},
				counterMetrics: map[string]int64{},
			},
			want: 1,
		}, {
			name: "Test 2",
			args: args{
				memStats:     &runtime.MemStats{},
				gaugeMetrics: map[string]float64{},
				counterMetrics: map[string]int64{
					"PollCount": 10,
				},
			},
			want: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			CollectData(test.args.memStats, test.args.gaugeMetrics, test.args.counterMetrics)
			assert.Equal(t, test.want, test.args.counterMetrics["PollCount"])
		})
	}
}
