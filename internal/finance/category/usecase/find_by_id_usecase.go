package usecase

import (
	"github.com/brunosilv96/simple_finance_api/internal/finance/category"
	entity "github.com/brunosilv96/simple_finance_api/internal/finance/category/entity"
	"github.com/brunosilv96/simple_finance_api/internal/shared"
)

type FindCategoryById struct {
	CategoryRepository category.CategoryRepository
}

func NewFindCategoryById(categoryRepository category.CategoryRepository) *FindCategoryById {
	return &FindCategoryById{
		CategoryRepository: categoryRepository,
	}
}

func (usecase *FindCategoryById) Execute(id string) (entity.Category, error) {
	if id == "" {
		return entity.Category{}, &shared.InputCannotBeNil{
			Input: "id",
		}
	}

	foundCategory, err := usecase.CategoryRepository.FindByID(id)
	if err != nil {
		return entity.Category{}, category.CategoryNotFound
	}

	return *foundCategory, nil
}
