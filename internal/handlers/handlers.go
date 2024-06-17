package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"metrics/internal/storage"
	"metrics/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

func GetMetricsHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Print("start GetMetricsHandler")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	gaugeMetrics, counterMetrics := storage.Storage.GetMetrics()
	html := utils.GetHTMLWithMetrics(gaugeMetrics, counterMetrics)
	fmt.Print("GetMetricsHandler")
	io.WriteString(res, html)
}

func GetMetricValueHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Print("start GetMetricValueHandler")
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")

	switch metricType {
	case "gauge":
		{
			val, ok := storage.Storage.GetGaugeMetric(metricName)
			if !ok {
				http.Error(res, "Unknown metric name", http.StatusNotFound)
				return
			}

			output := strconv.FormatFloat(val, 'f', 3, 64)
			output = strings.TrimRight(output, "0")
			output = strings.TrimRight(output, ".")

			fmt.Print("GetMetricValueHandler")
			io.WriteString(res, output)
		}
	case "counter":
		{
			val, ok := storage.Storage.GetCounterMetric(metricName)
			if !ok {
				http.Error(res, "Unknown metric name", http.StatusNotFound)
				return
			}
			fmt.Print("GetMetricValueHandler")
			io.WriteString(res, strconv.FormatInt(val, 10))
		}
	default:
		{
			fmt.Print("GetMetricValueHandler")
			http.Error(res, "Unknown metric type", http.StatusNotFound)
			return
		}
	}
}

func UpdateMetricHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Print("start UpdateMetricHandler")
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	metricValue := chi.URLParam(req, "value")

	switch metricType {
	case "gauge":
		{
			gaugeValue, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				fmt.Print("finish UpdateMetricHandler")
				http.Error(res, "Invalid metric value", http.StatusBadRequest)
				return
			}

			storage.Storage.SetGaugeMetric(metricName, gaugeValue)
		}
	case "counter":
		{
			counterValue, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				fmt.Print("finish UpdateMetricHandler")
				http.Error(res, "Invalid metric value", http.StatusBadRequest)
				return
			}

			storage.Storage.SetCounterMetric(metricName, counterValue)
		}
	default:
		{
			fmt.Print("finish UpdateMetricHandler")
			http.Error(res, "Invalid metric type", http.StatusBadRequest)
			return
		}
	}

	fmt.Print("finish UpdateMetricHandler")
	res.WriteHeader(http.StatusOK)
}
