package handlers

import (
	"metrics/internal/storage"
	"metrics/internal/utils"
	"net/http"
)

func GetMetricsHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	gaugeMetrics, counterMetrics := storage.Storage.GetMetrics()
	html := utils.GetHTMLWithMetrics(gaugeMetrics, counterMetrics)
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(html))
}
