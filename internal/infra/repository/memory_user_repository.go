package repository

import (
	"errors"

	entity "github.com/brunosilv96/simple_finance_api/internal/finance/user/entity"
)

type MemoryUser struct {
	users map[string]entity.User
}

func NewMemoryUser() *MemoryUser {
	return &MemoryUser{
		users: make(map[string]entity.User),
	}
}

func (repository *MemoryUser) Save(user *entity.User) error {
	repository.users[user.ID] = *user

	return nil
}

func (repository *MemoryUser) FindByID(userID string) (*entity.User, error) {
	user, exist := repository.users[userID]
	if !exist {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
