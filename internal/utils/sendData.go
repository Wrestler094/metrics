package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"metrics/internal/models"
	"net/http"
)

func sendGaugeMetric(server string, k string, v float64) {
	url := fmt.Sprintf("%s/update/", server)

	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	jsonEncoder := json.NewEncoder(gzipWriter)

	err := jsonEncoder.Encode(models.Metrics{
		ID:    k,
		MType: "gauge",
		Value: &v,
	})

	if err != nil {
		fmt.Println("Encoding or compression error:", err)
		return
	}

	if err := gzipWriter.Close(); err != nil {
		fmt.Println("Gzip writer close error:", err)
		return
	}

	resp, err := http.Post(url, "application/json", &buf)
	if err != nil {
		fmt.Printf("Metric %s sent with Error %s\n", k, err)
		return
	}

	fmt.Printf("Metrics %s sent with value %f\n", k, v)

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

	fmt.Printf("Metrics %s sent with value %d\n", k, v)

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
