package usecase

import (
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

func (usecase *FindUserByID) Execute(userID string) (entity.User, error) {
	if userID == "" {
		return entity.User{}, &shared.InputCannotBeNil{
			Input: "user id",
		}
	}

	foundUser, err := usecase.userRepository.FindByID(userID)
	if err != nil {
		return entity.User{}, user.UserNotFound
	}

	return *foundUser, nil
}
