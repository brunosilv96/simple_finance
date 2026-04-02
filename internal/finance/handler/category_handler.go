package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"simple_finance/internal/finance/dto"
	usecaseError "simple_finance/internal/finance/usecase"
	usecase "simple_finance/internal/finance/usecase/category"
)

type CategoryHandler struct {
	CreateCategoryUC usecase.CreateCategory
	FindAllCategoriesUC usecase.FindAllCategories
	FindCategoryByIdUC usecase.FindCategoryById
	DeleteCategoryUC usecase.DeleteCategory
}

func NewCategoryHandler(createCategoryUseCase usecase.CreateCategory, listCategoriesUseCase usecase.FindAllCategories, findCategoryByIdUC usecase.FindCategoryById, deleteCategoryUseCase usecase.DeleteCategory) *CategoryHandler {
	return &CategoryHandler{
		CreateCategoryUC: createCategoryUseCase,
		FindAllCategoriesUC: listCategoriesUseCase,
		FindCategoryByIdUC: findCategoryByIdUC,
		DeleteCategoryUC: deleteCategoryUseCase,
	}
}

func (handler *CategoryHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/categories", handler.Create)
    mux.HandleFunc("GET /api/v1/categories", handler.FindAll)
    mux.HandleFunc("GET /api/v1/categories/{id}", handler.FindByID)
    mux.HandleFunc("DELETE /api/v1/categories/{id}", handler.Delete)
}

func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCategoryRequest

	err := input.Bind(r.Body)
	if err != nil {
		dto.ErrorReponse(w, http.StatusBadRequest, "invalid request body")
		return
	}
	
	category, err := handler.CreateCategoryUC.Execute(
		input.Name,
		input.Description,
	)
	if err != nil {
		dto.ErrorReponse(w, http.StatusBadRequest, err.Error())
		return
	}

	categoryResponse := &dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
		CreatedAt: category.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	categoryResponse.Render(w)
}

func (handler *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.FindAllCategoriesUC.Execute()
	if err != nil {
		dto.ErrorReponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var categoriesResponse []dto.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, dto.CategoryResponse{
			ID: category.ID,
			Name: category.Name,
			Description: category.Description,
			CreatedAt: category.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoriesResponse)
}

func (handler *CategoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		dto.ErrorReponse(w, http.StatusBadRequest, "ID is required")
        return
    }

	category, err := handler.FindCategoryByIdUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, usecaseError.CategoryNotFound):
			dto.ErrorReponse(w, http.StatusNotFound, err.Error())
			return
		default:
			dto.ErrorReponse(w, http.StatusInternalServerError, "Server Internal Error")
			return

		}
	}

	categoryResponse := &dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
		CreatedAt: category.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	categoryResponse.Render(w)
}

func (handler *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		dto.ErrorReponse(w, http.StatusBadRequest, "ID is required")
        return
    }

	err := handler.DeleteCategoryUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, usecaseError.CategoryNotFound):
			dto.ErrorReponse(w, http.StatusNotFound, err.Error())
			return
		default:
			dto.ErrorReponse(w, http.StatusInternalServerError, "Server Internal Error")
			return

		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w)
}
