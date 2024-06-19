package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (bh *BaseHandler) postMetricHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	metricType := chi.URLParam(r, "type")
	metricName := chi.URLParam(r, "name")
	metricValue := chi.URLParam(r, "value")

	switch metricType {
	case "gauge":
		{
			gaugeValue, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, "Invalid metric value", http.StatusBadRequest)
				return
			}

			bh.Storage.SetGaugeMetric(metricName, gaugeValue)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}
	case "counter":
		{
			counterValue, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(w, "Invalid metric value", http.StatusBadRequest)
				return
			}

			bh.Storage.SetCounterMetric(metricName, counterValue)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}
	default:
		{
			http.Error(w, "Invalid metric type", http.StatusBadRequest)
			return
		}
	}
}
