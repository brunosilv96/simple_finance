package usecase

import (
	"context"
	"errors"
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

func (usecase *RegisterUser) Execute(ctx context.Context, name string) (*entity.User, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("http connection was cancel")
	default:
	}

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

	err := usecase.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
