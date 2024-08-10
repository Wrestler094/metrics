package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"metrics/internal/models"
	"net/http"
)

func sendGaugeMetric(server string, k string, v float64) {
	url := fmt.Sprintf("%s/update/", server)

	out, err := json.Marshal(models.Metrics{
		ID:    k,
		MType: "gauge",
		Value: &v,
	})

	if err != nil {
		fmt.Println("Serialisation error", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(out))
	if err != nil {
		fmt.Printf("Metric %s sent with Error %s\n", k, err)
		return
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Error closing response body:", closeErr)
		}
	}()
}

func sendCounterMetric(server string, k string, v int64) {
	url := fmt.Sprintf("%s/update/", server)

	out, err := json.Marshal(models.Metrics{
		ID:    k,
		MType: "counter",
		Delta: &v,
	})

	if err != nil {
		fmt.Println("Serialisation error", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(out))
	if err != nil {
		fmt.Printf("Metric %s sent with Error %s\n", k, err)
		return
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Error closing response body:", closeErr)
		}
	}()
}

func SendData(gaugeMetrics map[string]float64, counterMetrics map[string]int64, server string) {
	for k, v := range gaugeMetrics {
		sendGaugeMetric(server, k, v)
	}

	for k, v := range counterMetrics {
		sendCounterMetric(server, k, v)
	}
}
