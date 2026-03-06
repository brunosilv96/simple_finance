package finance

import (
	entities "simple_finance/internal/finance/entities"
	"time"

	"github.com/google/uuid"
)

func CreateUserExecute(name string) (*entities.User, error){
	if name == "" {
		return nil, &entities.InputCannotBeNil{
			Input: "name",
		}
	}

	user := &entities.User{
		ID: uuid.NewString(),
		Name: name,
		CreatedAt: time.Now(),
	}

	return user, nil
}
