package usecase

import (
	"context"
	"errors"

	"github.com/brunosilv96/simple_finance_api/internal/finance/category"
	entity "github.com/brunosilv96/simple_finance_api/internal/finance/category/entity"
)

type FindAllCategories struct {
	CategoryRepository category.CategoryRepository
}

func NewListCategories(categoryRepository category.CategoryRepository) *FindAllCategories {
	return &FindAllCategories{
		CategoryRepository: categoryRepository,
	}
}

func (usecase *FindAllCategories) Execute(ctx context.Context) ([]entity.Category, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("http connection was cancel")
	default:
	}

	categories, err := usecase.CategoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
