package usecase

import (
	"time"

	"github.com/brunosilv96/simple_finance_api/internal/finance/user"
	entity "github.com/brunosilv96/simple_finance_api/internal/finance/user/entity"
	"github.com/brunosilv96/simple_finance_api/internal/shared"

	"github.com/google/uuid"
)

type RegisterUser struct {
	userRepository user.UserRepository
}

func NewRegisterUser(repository user.UserRepository) *RegisterUser {
	return &RegisterUser{
		userRepository: repository,
	}
}

func (usecase *RegisterUser) Execute(name string) (*entity.User, error) {
	if name == "" {
		return nil, &shared.InputCannotBeNil{
			Input: "name",
		}
	}

	user := &entity.User{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := usecase.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
