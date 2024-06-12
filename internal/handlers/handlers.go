package handlers

import (
	"metrics/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func HandleUpdateMetric(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(req.URL.Path, "/")
	urlParts = urlParts[2:]

	if len(urlParts) < 2 || urlParts[1] == "" {
		http.Error(res, "Request without metric name", http.StatusNotFound)
		return
	}

	if len(urlParts) != 3 {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	metricType := strings.ToLower(urlParts[0])
	metricName := urlParts[1]
	metricValue := urlParts[2]

	switch metricType {
	case "gauge":
		{
			gaugeValue, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(res, "Invalid metric value", http.StatusBadRequest)
				return
			}

			storage.Storage.ChangeGauge(metricName, gaugeValue)
		}
	case "counter":
		{
			counterValue, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(res, "Invalid metric value", http.StatusBadRequest)
				return
			}

			storage.Storage.IncreaseCounter(metricName, counterValue)
		}
	default:
		{
			http.Error(res, "Invalid metric type", http.StatusBadRequest)
			return
		}
	}

	res.WriteHeader(http.StatusOK)
}
