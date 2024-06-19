package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
)

func (bh *BaseHandler) getMetricValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	metricType := chi.URLParam(r, "type")
	metricName := chi.URLParam(r, "name")

	switch metricType {
	case "gauge":
		{
			val, ok := bh.Storage.GetGaugeMetric(metricName)
			if !ok {
				http.Error(w, "Unknown metric name", http.StatusNotFound)
				return
			}

			// output := strconv.FormatFloat(val, 'f', 3, 64)
			// output = strings.TrimRight(output, "0")
			// output = strings.TrimRight(output, ".")
			output := strings.Trim(fmt.Sprintf("%.5f", val), "0 .")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(output))
		}
	case "counter":
		{
			val, ok := bh.Storage.GetCounterMetric(metricName)
			if !ok {
				http.Error(w, "Unknown metric name", http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(strconv.FormatInt(val, 10)))
		}
	default:
		{
			http.Error(w, "Unknown metric type", http.StatusNotFound)
			return
		}
	}
}
