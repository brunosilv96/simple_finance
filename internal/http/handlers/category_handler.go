package http

import (
	"errors"
	"net/http"

	domain "github.com/brunosilv96/simple_finance_api/internal/finance/category"
	usecase "github.com/brunosilv96/simple_finance_api/internal/finance/category/usecase"
	"github.com/brunosilv96/simple_finance_api/internal/http/dto"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CreateCategoryUC    usecase.CreateCategory
	FindAllCategoriesUC usecase.FindAllCategories
	FindCategoryByIdUC  usecase.FindCategoryById
	DeleteCategoryUC    usecase.DeleteCategory
}

func NewCategoryHandler(repository domain.CategoryRepository) *CategoryHandler {
	createCategoryUseCase := usecase.NewCreateCategory(repository)
	listCategoriesUseCase := usecase.NewListCategories(repository)
	categoryByIdUseCase := usecase.NewFindCategoryById(repository)
	deleteCategoryUseCase := usecase.NewDeleteCategory(repository)

	return &CategoryHandler{
		CreateCategoryUC:    *createCategoryUseCase,
		FindAllCategoriesUC: *listCategoriesUseCase,
		FindCategoryByIdUC:  *categoryByIdUseCase,
		DeleteCategoryUC:    *deleteCategoryUseCase,
	}
}

func (handler *CategoryHandler) Create(c *gin.Context) {
	var input dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	category, err := handler.CreateCategoryUC.Execute(
		input.Name,
		input.Description,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	categoryResponse := &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
	}

	c.JSON(http.StatusCreated, categoryResponse)
}

func (handler *CategoryHandler) FindAll(c *gin.Context) {
	categories, err := handler.FindAllCategoriesUC.Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var categoriesResponse []dto.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, categoriesResponse)
}

func (handler *CategoryHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	category, err := handler.FindCategoryByIdUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.CategoryNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return

		}
	}

	categoryResponse := &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
	}

	c.JSON(http.StatusOK, categoryResponse)
}

func (handler *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	err := handler.DeleteCategoryUC.Execute(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.CategoryNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return

		}
	}

	c.Status(http.StatusNoContent)
}
