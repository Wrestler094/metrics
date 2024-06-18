package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (bh *BaseHandler) postMetricHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	r.Close = true

	metricType := chi.URLParam(r, "type")
	metricName := chi.URLParam(r, "name")
	metricValue := chi.URLParam(r, "value")

	switch metricType {
	case "gauge":
		{
			fmt.Printf("1")
			gaugeValue, err := strconv.ParseFloat(metricValue, 64)
			fmt.Printf("2")
			if err != nil {
				fmt.Printf("3")
				http.Error(w, "Invalid metric value", http.StatusBadRequest)
				return
			}

			fmt.Printf("4")
			bh.storage.SetGaugeMetric(metricName, gaugeValue)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			fmt.Printf("5")
		}
	case "counter":
		{
			counterValue, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(w, "Invalid metric value", http.StatusBadRequest)
				return
			}

			bh.storage.SetCounterMetric(metricName, counterValue)
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
