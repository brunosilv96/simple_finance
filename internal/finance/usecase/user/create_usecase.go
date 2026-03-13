package usecase

import (
	"simple_finance/internal/finance/entity"
	"time"

	errors "simple_finance/internal/finance/usecase"

	"github.com/google/uuid"
)

type CreateUser struct {}

func (usecase *CreateUser) Execute(name string) (*entity.User, error){
	if name == "" {
		return nil, &errors.InputCannotBeNil{
			Input: "name",
		}
	}

	user := &entity.User{
		ID: uuid.NewString(),
		Name: name,
		CreatedAt: time.Now(),
	}

	return user, nil
}
