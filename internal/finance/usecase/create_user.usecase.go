package usecase

import (
	"simple_finance/internal/finance/entity"
	"time"

	"github.com/google/uuid"
)

type CreateUserUseCase struct {}

func (usecase *CreateUserUseCase) Execute(name string) (*entity.User, error){
	if name == "" {
		return nil, &entity.InputCannotBeNil{
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
