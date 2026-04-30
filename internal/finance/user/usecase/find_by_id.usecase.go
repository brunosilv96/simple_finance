package usecase

import (
	"context"
	"errors"

	"github.com/brunosilv96/simple_finance_api/internal/finance/user"
	entity "github.com/brunosilv96/simple_finance_api/internal/finance/user/entity"
	"github.com/brunosilv96/simple_finance_api/internal/shared"
)

type FindUserByID struct {
	userRepository user.UserRepository
}

func NewFindUserByID(repository user.UserRepository) *FindUserByID {
	return &FindUserByID{
		userRepository: repository,
	}
}

func (usecase *FindUserByID) Execute(ctx context.Context, userID string) (entity.User, error) {
	select {
	case <-ctx.Done():
		return entity.User{}, errors.New("http connection was cancel")
	default:
	}

	if userID == "" {
		return entity.User{}, &shared.InputCannotBeNil{
			Input: "user id",
		}
	}

	foundUser, err := usecase.userRepository.FindByID(ctx, userID)
	if err != nil {
		return entity.User{}, user.UserNotFound
	}

	return *foundUser, nil
}
