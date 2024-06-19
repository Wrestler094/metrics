package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendData(t *testing.T) {
	type args struct {
		gaugeMetrics   map[string]float64
		counterMetrics map[string]int64
		server         string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				gaugeMetrics: map[string]float64{
					"RandomValue": 10.5,
				},
				counterMetrics: map[string]int64{},
			},
		},
		{
			name: "Test 2",
			args: args{
				gaugeMetrics: map[string]float64{},
				counterMetrics: map[string]int64{
					"PollCount": 20,
				},
			},
		},
		{
			name: "Test 3",
			args: args{
				gaugeMetrics: map[string]float64{
					"RandomValue": 10.5,
					"TotalAlloc":  10.5432123,
					"Sys":         10,
				},
				counterMetrics: map[string]int64{
					"PollCount": 20,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "POST", r.Method, "Expected a POST request")
				assert.Equal(t, "text/plain", r.Header.Get("Content-Type"), "Expected a text/plain Content-Type")
				assert.Equal(t, 5, len(strings.Split(r.URL.Path, "/")))
				assert.Contains(t, strings.Split(r.URL.Path, "/"), "update")
			}))
			defer server.Close()

			SendData(test.args.gaugeMetrics, test.args.counterMetrics, server.URL)
		})
	}
}
