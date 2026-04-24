package usecase

import (
	"time"

	entity "github.com/brunosilv96/simple_finance_api/internal/finance/debit/entity"
	"github.com/brunosilv96/simple_finance_api/internal/shared"

	"github.com/google/uuid"
)

type RegisterDebit struct{}

func (usecase *RegisterDebit) Execute(categoryID, userID, title, description string, date time.Time, value float64) (*entity.Debit, error) {
	if categoryID == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "category id",
		}
	}

	if userID == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "user id",
		}
	}

	if title == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "title",
		}
	}

	if description == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "description",
		}
	}

	if value <= 0 {
		return nil, &shared.InputCannotBeNil{
			Input: "value",
		}
	}

	debit := &entity.Debit{
		ID:          uuid.NewString(),
		UserID:      userID,
		CategoryID:  categoryID,
		Title:       title,
		Description: description,
		Value:       value,
		Date:        date,
	}

	return debit, nil
}
