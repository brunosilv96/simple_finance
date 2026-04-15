package usecase

import (
	"github.com/brunosilv96/simple_finance_api/internal/finance/entity"
	errors "github.com/brunosilv96/simple_finance_api/internal/finance/usecase"
)

type FindUserByID struct {
	userRespository UserRepository
}

func NewFindUserByID(repository UserRepository) *FindUserByID {
	return &FindUserByID{
		userRespository: repository,
	}
}

func (usecase *FindUserByID) Execute(userID string) (entity.User, error) {
	if userID == "" {
		return entity.User{}, &errors.InputCannotBeNil{
			Input: "user id",
		}
	}

	user, err := usecase.userRespository.FindByID(userID)
	if err != nil {
		return entity.User{}, errors.UserNotFound
	}

	return *user, nil
}
