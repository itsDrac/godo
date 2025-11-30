package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/itsDrac/godo/internal/service"
)

type Handler interface {
	Router() http.Handler
	Mount()
}

type ChiHandler struct {
	router *chi.Mux
	ser service.Servicer

}

func (h *ChiHandler) Router() http.Handler {
	return h.router
}

func NewChiHandler(ser service.Servicer) Handler {
	r := chi.NewRouter()
	
	return &ChiHandler{
		router: r,
		ser:      ser,
	}
}

func (h *ChiHandler) Mount() {
	r := h.router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/api", func(r chi.Router) {
		r.Post("users", h.CreateUser)
	})
}