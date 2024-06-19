package handlers

import (
	"metrics/internal/utils"
	"net/http"
)

func (bh *BaseHandler) getMetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	gaugeMetrics, counterMetrics := bh.Storage.GetMetrics()
	html := utils.GetHTMLWithMetrics(gaugeMetrics, counterMetrics)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
