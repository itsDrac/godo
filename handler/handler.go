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
	ser    service.Servicer
}

func (h *ChiHandler) Router() http.Handler {
	return h.router
}

func NewChiHandler(ser service.Servicer) Handler {
	r := chi.NewRouter()

	return &ChiHandler{
		router: r,
		ser:    ser,
	}
}

func (h *ChiHandler) Mount() {
	r := h.router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Simple CORS middleware to allow the frontend served from Live Server (http://localhost:5500)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			// Allow both localhost and 127.0.0.1 served pages on port 5500 during development
			if origin == "http://localhost:5500" || origin == "http://127.0.0.1:5500" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	r.Route("/api", func(r chi.Router) {
		r.Post("/users", h.CreateUser)
		r.Post("/login", h.Login)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./web"))))
}
