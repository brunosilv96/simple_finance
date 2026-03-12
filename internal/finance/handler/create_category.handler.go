package handler

import (
	"encoding/json"
	"net/http"
	"simple_finance/internal/finance/usecase"
)

type CreateCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryHandler struct {
	CreateCategoryUseCase usecase.CreateCategoryUseCase
}

func (handler *CreateCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input CreateCategoryInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	
	category, err := handler.CreateCategoryUseCase.Execute(
		input.Name,
		input.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
