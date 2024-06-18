package handlers

import (
	"github.com/go-chi/chi/v5"
	"metrics/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

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

			res.WriteHeader(http.StatusOK)
			res.Write([]byte(output))
		}
	case "counter":
		{
			val, ok := storage.Storage.GetCounterMetric(metricName)
			if !ok {
				http.Error(res, "Unknown metric name", http.StatusNotFound)
				return
			}

			res.WriteHeader(http.StatusOK)
			res.Write([]byte(strconv.FormatInt(val, 10)))
		}
	default:
		{
			http.Error(res, "Unknown metric type", http.StatusNotFound)
			return
		}
	}
}
