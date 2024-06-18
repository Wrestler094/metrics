package handlers

import (
	"metrics/internal/storage"

	"github.com/go-chi/chi/v5"
)

type BaseHandler struct {
	Storage storage.Repository
}

func NewBaseHandler(storage storage.Repository) *BaseHandler {
	return &BaseHandler{Storage: storage}
}

func (bh *BaseHandler) Router() *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middleware.Logger)
	// router.Use(middleware.Recoverer)

	router.Get("/", bh.getMetricsHandler)
	router.Get("/value/{type}/{name}", bh.getMetricValueHandler)
	router.Post("/update/{type}/{name}/{value}", bh.postMetricHandler)

	return router
}
