package utils

import (
	"fmt"
	"net/http"
)

func SendData(gaugeMetrics map[string]float64, counterMetrics map[string]int64, server string) {
	for k, v := range gaugeMetrics {
		url := fmt.Sprintf("%s/update/gauge/%s/%f", server, k, v)
		resp, err := http.Post(url, "text/plain", nil)
		if err != nil {
			fmt.Printf("Metric %s sent with Error %s\n", k, err)
		}
		defer resp.Body.Close()
	}

	for k, v := range counterMetrics {
		url := fmt.Sprintf("%s/update/counter/%s/%d", server, k, v)
		resp, err := http.Post(url, "text/plain", nil)
		if err != nil {
			fmt.Printf("Metric %s sent with Error %s\n", k, err)
		}
		defer resp.Body.Close()
	}
}
