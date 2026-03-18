package dto

import (
	"encoding/json"
	"net/http"
)

type APIErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func ErrorReponse(w http.ResponseWriter, code int, message string) error {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

	dto := &APIErrorResponse{
		Code: code,
		Message: message,
	}

	return json.NewEncoder(w).Encode(dto)
}