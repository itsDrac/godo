package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
		log.Printf("CreateUser error: %v", err)
		le := strings.ToLower(err.Error())
		if strings.Contains(le, "duplicate") || strings.Contains(le, "unique") {
			http.Error(w, "Account with that email already exists", http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
