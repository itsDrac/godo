package handler

import (
	"encoding/json"
	"net/http"

	"github.com/itsDrac/godo/internal/service"
)

func (h *ChiHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req service.LoginParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	res, err := h.authService.Login(r.Context(), req)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
