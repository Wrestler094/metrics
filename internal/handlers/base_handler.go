package handlers

import (
	"github.com/go-chi/chi/v5"
	"metrics/internal/compressor"
	"metrics/internal/logger"
	"metrics/internal/storage"
)

type BaseHandler struct {
	Storage storage.Repository
}

func NewBaseHandler(storage storage.Repository) *BaseHandler {
	return &BaseHandler{Storage: storage}
}

func (bh *BaseHandler) Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(compressor.GzipMiddleware)
	router.Use(logger.WithLogging)

	router.Get("/", bh.getMetricsHandler)
	router.Get("/value/{type}/{name}", bh.getMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", bh.postMetricValueHandler)
	router.Post("/update", bh.postUpdateValueHandler)
	router.Post("/update/", bh.postUpdateValueHandler)
	router.Post("/value", bh.postGetValueHandler)
	router.Post("/value/", bh.postGetValueHandler)

	return router
}
