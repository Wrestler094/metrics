package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"metrics/internal/storage"
	"metrics/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

func GetMetricsHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	gaugeMetrics, counterMetrics := storage.Storage.GetMetrics()
	html := utils.GetHTMLWithMetrics(gaugeMetrics, counterMetrics)
	_, err := io.WriteString(res, html)
	if err != nil {
		log.Fatalf("io.WriteString error %v", err)
	}
}

func GetMetricValueHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")

	switch metricType {
	case "gauge":
		{
			val, ok := storage.Storage.GetGaugeMetric(metricName)
			if !ok {
				http.Error(res, "Unknown metric name", http.StatusNotFound)
				return
			}

			output := strconv.FormatFloat(val, 'f', 3, 64)
			output = strings.TrimRight(output, "0")
			output = strings.TrimRight(output, ".")

			_, err := io.WriteString(res, output)
			if err != nil {
				log.Fatalf("io.WriteString error %v", err)
			}
		}
	case "counter":
		{
			val, ok := storage.Storage.GetCounterMetric(metricName)
			if !ok {
				http.Error(res, "Unknown metric name", http.StatusNotFound)
				return
			}

			_, err := io.WriteString(res, strconv.FormatInt(val, 10))
			if err != nil {
				log.Fatalf("io.WriteString error %v", err)
			}
		}
	default:
		{
			http.Error(res, "Unknown metric type", http.StatusNotFound)
			return
		}
	}
}

func UpdateMetricHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	metricValue := chi.URLParam(req, "value")

	switch metricType {
	case "gauge":
		{
			gaugeValue, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(res, "Invalid metric value", http.StatusBadRequest)
				return
			}

			storage.Storage.SetGaugeMetric(metricName, gaugeValue)
		}
	case "counter":
		{
			counterValue, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(res, "Invalid metric value", http.StatusBadRequest)
				return
			}

			storage.Storage.SetCounterMetric(metricName, counterValue)
		}
	default:
		{
			http.Error(res, "Invalid metric type", http.StatusBadRequest)
			return
		}
	}

	res.WriteHeader(http.StatusOK)
}
