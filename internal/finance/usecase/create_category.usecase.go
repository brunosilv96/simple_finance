package usecase

import (
	"simple_finance/internal/finance/entity"
	"time"

	"github.com/google/uuid"
)

type CreateCategoryUseCase struct {}

func (usecase *CreateCategoryUseCase) Execute(name, description string) (*entity.Category, error){
	if name == "" {
		return nil, &entity.InputCannotBeNil{
			Input: "name",
		}
	}

	category := &entity.Category{
		ID: uuid.NewString(),
		Name: name,
		Description: description,
		CreatedAt: time.Now(),
	}

	return category, nil
}
