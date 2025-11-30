package handler

import (
	"encoding/json"
	"net/http"

	"github.com/itsDrac/godo/internal/service"
)


func (h *ChiHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var req service.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.ser.CreateUser(r.Context(), req)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}