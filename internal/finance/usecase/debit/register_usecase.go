package usecase

import (
	"time"

	"github.com/brunosilv96/simple_finance_api/internal/finance/entity"

	errors "github.com/brunosilv96/simple_finance_api/internal/finance/usecase"

	"github.com/google/uuid"
)

type RegisterDebit struct{}

func (usecase *RegisterDebit) Execute(categoryID, userID, title, description string, date time.Time, value float64) (*entity.Debit, error) {
	if categoryID == "" {
		return nil, &errors.InputCannotBeNil{
			Input: "category id",
		}
	}

	if userID == "" {
		return nil, &errors.InputCannotBeNil{
			Input: "user id",
		}
	}

	if title == "" {
		return nil, &errors.InputCannotBeNil{
			Input: "title",
		}
	}

	if description == "" {
		return nil, &errors.InputCannotBeNil{
			Input: "description",
		}
	}

	if value <= 0 {
		return nil, &errors.InputCannotBeNil{
			Input: "value",
		}
	}

	debit := &entity.Debit{
		ID:          uuid.NewString(),
		UserID:      userID,
		CategoryID:  categoryID,
		Title:       title,
		Desctiption: description,
		Value:       value,
		Date:        date,
	}

	return debit, nil
}
