package handlers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"metrics/internal/logger"
	"metrics/internal/models"
	"net/http"
)

func (bh *BaseHandler) postGetValueHandler(w http.ResponseWriter, r *http.Request) {
	var metric models.Metrics
	logger.Log.Info("Start")

	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.Info("metric", zap.Any("metric", metric))
	w.Header().Set("Content-Type", "application/json")

	switch metric.MType {
	case "gauge":
		{
			logger.Log.Info("gauge")
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
			logger.Log.Info("counter")
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
