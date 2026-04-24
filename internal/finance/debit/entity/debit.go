package entity

import (
	"time"

	category "github.com/brunosilv96/simple_finance_api/internal/finance/category/entity"
	user "github.com/brunosilv96/simple_finance_api/internal/finance/user/entity"
	"github.com/brunosilv96/simple_finance_api/internal/shared"

	"github.com/google/uuid"
)

type Debit struct {
	ID          string
	CategoryID  string
	UserID      string
	Title       string
	Description string
	Value       float64
	Date        time.Time
}

func NewDebit(category category.Category, user user.User, title, description string, date time.Time, value float64) (*Debit, error) {
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

	debit := &Debit{
		ID:          uuid.NewString(),
		UserID:      user.ID,
		CategoryID:  category.ID,
		Title:       title,
		Description: description,
		Value:       value,
		Date:        date,
	}

	return debit, nil
}
