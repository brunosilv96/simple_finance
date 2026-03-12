package entity

import (
	"time"

	"github.com/google/uuid"
)

type Debit struct {
	ID string
	CategoryID string
	UserID string
	Title string
	Desctiption string
	Value float64
	Date time.Time
}

func NewDebit(category Category, user User, title, description string, date time.Time, value float64) (*Debit, error) {
	if title == "" {
		return nil, &InputCannotBeNil{
			Input: "title",
		}
	}

	if description == "" {
		return nil, &InputCannotBeNil{
			Input: "description",
		}
	}

	if value <= 0 {
		return nil, &InputCannotBeNil{
			Input: "value",
		}
	}

	debit := &Debit{
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