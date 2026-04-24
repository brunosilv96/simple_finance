package entity

import (
	"time"

	"github.com/brunosilv96/simple_finance_api/internal/shared"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "name",
		}
	}

	user := &User{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return user, nil
}
