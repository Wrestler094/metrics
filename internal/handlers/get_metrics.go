package handlers

import (
	"go.uber.org/zap"
	"metrics/internal/logger"
	"metrics/internal/utils"
	"net/http"
)

func (bh *BaseHandler) getMetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	gaugeMetrics, counterMetrics := bh.Storage.GetMetrics()
	html := utils.GetHTMLWithMetrics(gaugeMetrics, counterMetrics)
	_, err := w.Write([]byte(html))
	if err != nil {
		logger.Log.Info("Failed to write response", zap.Error(err))
	}
}
