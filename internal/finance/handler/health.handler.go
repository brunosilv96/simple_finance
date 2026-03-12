package handler

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string {
		"status": "ok",
	}

	// Monta header do response
	w.Header().Set("Context-Type", "application/json")

	// 1. Converte o map response para json
	// 2. Responde diereto no Response Writer
	json.NewEncoder(w).Encode(response)
}