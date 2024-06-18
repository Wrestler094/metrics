package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func sendGaugeMetric(server string, k string, v float64) {
	url := fmt.Sprintf("%s/update/gauge/%s/3.333", server, k, strconv.FormatFloat(v, 'f', -1, 64))
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		fmt.Printf("Metric %s sent with Error %s\n", k, err)
	} else {
		defer resp.Body.Close()
	}
}

func sendCounterMetric(server string, k string, v int64) {
	url := fmt.Sprintf("%s/update/counter/%s/%d", server, k, v)
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		fmt.Printf("Metric %s sent with Error %s\n", k, err)
	} else {
		defer resp.Body.Close()
	}
}

func SendData(gaugeMetrics map[string]float64, counterMetrics map[string]int64, server string) {
	for k, v := range gaugeMetrics {
		sendGaugeMetric(server, k, v)
	}

	for k, v := range counterMetrics {
		sendCounterMetric(server, k, v)
	}
}
