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

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")

	client := &http.Client{}
	resp, err := client.Do(req)
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

	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	jsonEncoder := json.NewEncoder(gzipWriter)

	err := jsonEncoder.Encode(models.Metrics{
		ID:    k,
		MType: "counter",
		Delta: &v,
	})

	if err != nil {
		fmt.Println("Encoding or compression error:", err)
		return
	}

	if err := gzipWriter.Close(); err != nil {
		fmt.Println("Gzip writer close error:", err)
		return
	}

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")

	client := &http.Client{}
	resp, err := client.Do(req)
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
