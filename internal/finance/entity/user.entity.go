package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID string
	Name string
	CreatedAt time.Time
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, &InputCannotBeNil{
			Input: "name",
		}
	}

	user := &User{
		ID: uuid.NewString(),
		Name: name,
		CreatedAt: time.Now(),
	}

	return user, nil
}