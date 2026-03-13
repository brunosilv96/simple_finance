package handler

import (
	"encoding/json"
	"net/http"
	"simple_finance/internal/finance/dto"
	usecase "simple_finance/internal/finance/usecase/category"
)

type CategoryHandler struct {
	CreateCategoryUC usecase.CreateCategory
	FindAllCategoriesUC usecase.FindAllCategories
	FindCategoryByIdUC usecase.FindCategoryById
}

func NewCategoryHandler(createCategoryUseCase usecase.CreateCategory, listCategoriesUseCase usecase.FindAllCategories, findCategoryByIdUC usecase.FindCategoryById) *CategoryHandler {
	return &CategoryHandler{
		CreateCategoryUC: createCategoryUseCase,
		FindAllCategoriesUC: listCategoriesUseCase,
		FindCategoryByIdUC: findCategoryByIdUC,
	}
}

func (handler *CategoryHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/categories", handler.Create)
    mux.HandleFunc("GET /api/v1/categories", handler.FindAll)
    mux.HandleFunc("GET /api/v1/categories/{id}", handler.FindByID)
}

func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCategoryRequest

	err := input.Bind(r.Body)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	
	category, err := handler.CreateCategoryUC.Execute(
		input.Name,
		input.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func (handler *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.FindAllCategoriesUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (handler *CategoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }

	category, err := handler.FindCategoryByIdUC.Execute(id)
	if err != nil {
        if err.Error() == "category not found" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, "Erro interno", http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
