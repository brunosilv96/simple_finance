package usecase

import (
	"simple_finance/internal/finance/entity"
	"time"

	"github.com/google/uuid"
)

type CreateDebitUseCase struct {}

func (usecase *CreateDebitUseCase) Execute(category entity.Category, user entity.User, title, description string, date time.Time, value float64) (*entity.Debit, error){
	if title == "" {
		return nil, &entity.InputCannotBeNil{
			Input: "title",
		}
	}

	if description == "" {
		return nil, &entity.InputCannotBeNil{
			Input: "description",
		}
	}

	if value <= 0 {
		return nil, &entity.InputCannotBeNil{
			Input: "value",
		}
	}

	debit := &entity.Debit{
		ID: uuid.NewString(),
		UserID: user.ID,
		CategoryID: category.ID,
		Title: title,
		Desctiption: description,
		Value: value,
		Date: date,
	}

	return debit, nil
}