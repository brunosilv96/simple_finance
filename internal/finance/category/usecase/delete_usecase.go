package usecase

import (
	"context"
	"errors"

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

func (usecase *DeleteCategory) Execute(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return errors.New("http connection was cancel")
	default:
	}

	if id == "" {
		return &shared.InputCannotBeNil{
			Input: "id",
		}
	}

	err := usecase.CategoryRepository.Delete(ctx, id)
	if err != nil {
		return category.CategoryNotFound
	}

	return nil
}
