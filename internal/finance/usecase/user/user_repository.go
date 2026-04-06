package usecase

import "simple_finance/internal/finance/entity"

type UserRepository interface {
	Save(*entity.User) error
}