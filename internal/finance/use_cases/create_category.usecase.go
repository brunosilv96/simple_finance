package finance

import (
	entities "simple_finance/internal/finance/entities"

	"github.com/google/uuid"
)

func CreateCategoryExecute(name, description string) (*entities.Category, error){
	if name == "" {
		return nil, &entities.InputCannotBeNil{
			Input: "name",
		}
	}

	category := &entities.Category{
		ID: uuid.NewString(),
		Name: name,
		Description: description,
	}

	return category, nil
}
