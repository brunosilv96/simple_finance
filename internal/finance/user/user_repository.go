package user

import (
	entity "github.com/brunosilv96/simple_finance_api/internal/finance/user/entity"
)

type UserRepository interface {
	Save(*entity.User) error
	FindByID(userID string) (*entity.User, error)
}
