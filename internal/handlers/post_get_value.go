package handlers

import (
	"encoding/json"
	"fmt"
	"metrics/internal/models"
	"net/http"
)

func (bh *BaseHandler) postGetValueHandler(w http.ResponseWriter, r *http.Request) {
	var metric models.Metrics

	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch metric.MType {
	case "gauge":
		{
			gaugeValue, ok := bh.Storage.GetGaugeMetric(metric.ID)
			if !ok {
				http.Error(w, "Unknown metric name", http.StatusNotFound)
				return
			}

			resp := models.Metrics{
				ID:    metric.ID,
				MType: metric.MType,
				Value: &gaugeValue,
			}

			enc := json.NewEncoder(w)
			if err := enc.Encode(resp); err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}
		}
	case "counter":
		{
			counterValue, ok := bh.Storage.GetCounterMetric(metric.ID)
			if !ok {
				http.Error(w, "Unknown metric name", http.StatusNotFound)
				return
			}

			resp := models.Metrics{
				ID:    metric.ID,
				MType: metric.MType,
				Delta: &counterValue,
			}

			enc := json.NewEncoder(w)
			if err := enc.Encode(resp); err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}
		}
	default:
		{
			http.Error(w, "Unknown metric type", http.StatusBadRequest)
			return
		}
	}
}
