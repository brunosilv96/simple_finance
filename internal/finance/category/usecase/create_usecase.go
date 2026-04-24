package usecase

import (
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

func (usecase *CreateCategory) Execute(name, description string) (*entity.Category, error) {
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

	err := usecase.categoryRepository.Save(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
