package repository

import "github.com/brunosilv96/simple_finance_api/internal/finance/entity"

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
