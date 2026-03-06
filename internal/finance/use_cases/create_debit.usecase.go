package finance

import (
	entities "simple_finance/internal/finance/entities"
	"time"

	"github.com/google/uuid"
)

func CreateNewDebit(category entities.Category, user entities.User, title, description string, date time.Time, value float64) (*entities.Debit, error){
	if title == "" {
		return nil, &entities.InputCannotBeNil{
			Input: "title",
		}
	}

	if description == "" {
		return nil, &entities.InputCannotBeNil{
			Input: "description",
		}
	}

	if value <= 0 {
		return nil, &entities.InputCannotBeNil{
			Input: "value",
		}
	}

	debit := &entities.Debit{
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