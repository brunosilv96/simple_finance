package usecase

import (
	"context"
	"errors"

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

func (usecase *FindCategoryById) Execute(ctx context.Context, id string) (entity.Category, error) {
	select {
	case <-ctx.Done():
		return entity.Category{}, errors.New("http connection was cancel")
	default:
	}

	if id == "" {
		return entity.Category{}, &shared.InputCannotBeNil{
			Input: "id",
		}
	}

	foundCategory, err := usecase.CategoryRepository.FindByID(ctx, id)
	if err != nil {
		return entity.Category{}, category.CategoryNotFound
	}

	return *foundCategory, nil
}
