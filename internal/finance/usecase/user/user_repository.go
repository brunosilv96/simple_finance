package usecase

import "github.com/brunosilv96/simple_finance_api/internal/finance/entity"

type UserRepository interface {
	Save(*entity.User) error
}
