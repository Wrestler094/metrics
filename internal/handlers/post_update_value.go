package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"metrics/internal/logger"
	"metrics/internal/models"
	"net/http"
)

func (bh *BaseHandler) postUpdateValueHandler(w http.ResponseWriter, r *http.Request) {
	var metric models.Metrics

	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.Info("metric", zap.Any("metric", metric))
	w.Header().Set("Content-Type", "application/json")

	switch metric.MType {
	case "gauge":
		{
			if metric.Value == nil {
				http.Error(w, "Value doesn't exist in gauge metric", http.StatusBadRequest)
				return
			}

			bh.Storage.SetGaugeMetric(metric.ID, *metric.Value)
			resp := models.Metrics{
				ID:    metric.ID,
				MType: metric.MType,
				Value: metric.Value,
			}

			enc := json.NewEncoder(w)
			if err := enc.Encode(resp); err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}
		}
	case "counter":
		{
			if metric.Delta == nil {
				http.Error(w, "Delta doesn't exist in counter metric", http.StatusBadRequest)
				return
			}

			bh.Storage.SetCounterMetric(metric.ID, *metric.Delta)
			resp := models.Metrics{
				ID:    metric.ID,
				MType: metric.MType,
				Delta: metric.Delta,
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
