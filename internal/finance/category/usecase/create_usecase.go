package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/brunosilv96/simple_finance_api/internal/finance/category"
	entity "github.com/brunosilv96/simple_finance_api/internal/finance/category/entity"
	"github.com/brunosilv96/simple_finance_api/internal/shared"
	"github.com/google/uuid"
)

type CreateCategory struct {
	categoryRepository category.CategoryRepository
}

func NewCreateCategory(categoryRepository category.CategoryRepository) *CreateCategory {
	return &CreateCategory{
		categoryRepository: categoryRepository,
	}
}

func (usecase *CreateCategory) Execute(ctx context.Context, name, description string) (*entity.Category, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("http connection was cancel")
	default:
	}

	if name == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "name",
		}
	}

	category := &entity.Category{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}

	err := usecase.categoryRepository.Save(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
