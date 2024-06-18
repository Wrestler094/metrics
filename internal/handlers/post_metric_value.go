package handlers

import (
	"github.com/go-chi/chi/v5"
	"metrics/internal/storage"
	"net/http"
	"strconv"
)

func PostMetricHandler(res http.ResponseWriter, req *http.Request) {
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
	res.Write([]byte("OK"))
}
