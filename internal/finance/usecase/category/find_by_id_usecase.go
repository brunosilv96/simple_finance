package usecase

import (
	"github.com/brunosilv96/simple_finance_api/internal/finance/entity"
	errors "github.com/brunosilv96/simple_finance_api/internal/finance/usecase"
)

type FindCategoryById struct {
	CategoryRepository CategoryRepository
}

func NewFindCategoryById(categoryRepository CategoryRepository) *FindCategoryById {
	return &FindCategoryById{
		CategoryRepository: categoryRepository,
	}
}

func (usecase *FindCategoryById) Execute(id string) (entity.Category, error) {
	if id == "" {
		return entity.Category{}, &errors.InputCannotBeNil{
			Input: "id",
		}
	}

	category, err := usecase.CategoryRepository.FindByID(id)
	if err != nil {
		return entity.Category{}, errors.CategoryNotFound
	}

	return *category, nil
}
