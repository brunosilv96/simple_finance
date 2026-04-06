package usecase

import (
	"simple_finance/internal/finance/entity"
	"time"

	errors "simple_finance/internal/finance/usecase"

	"github.com/google/uuid"
)

type CreateCategory struct {
	categoryRepository CategoryRepository
}

func NewCreateCategory(categoryRepository CategoryRepository) *CreateCategory {
	return &CreateCategory{
		categoryRepository: categoryRepository,
	}
}

func (usecase *CreateCategory) Execute(name, description string) (*entity.Category, error){
	if name == "" {
		return nil, &errors.InputCannotBeNil{
			Input: "name",
		}
	}

	category := &entity.Category{
		ID: uuid.NewString(),
		Name: name,
		Description: description,
		CreatedAt: time.Now(),
	}

	err := usecase.categoryRepository.Save(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
