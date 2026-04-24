package usecase

import (
	"github.com/brunosilv96/simple_finance_api/internal/finance/category"
	"github.com/brunosilv96/simple_finance_api/internal/shared"
)

type DeleteCategory struct {
	CategoryRepository category.CategoryRepository
}

func NewDeleteCategory(categoryRepository category.CategoryRepository) *DeleteCategory {
	return &DeleteCategory{
		CategoryRepository: categoryRepository,
	}
}

func (usecase *DeleteCategory) Execute(id string) error {
	if id == "" {
		return &shared.InputCannotBeNil{
			Input: "id",
		}
	}

	err := usecase.CategoryRepository.Delete(id)
	if err != nil {
		return category.CategoryNotFound
	}

	return nil
}
