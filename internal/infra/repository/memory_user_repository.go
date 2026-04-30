package repository

import (
	"context"
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

func (repository *MemoryUser) Save(ctx context.Context, user *entity.User) error {
	select {
	case <-ctx.Done():
		return errors.New("http connection was cancel")
	default:
	}

	repository.users[user.ID] = *user

	return nil
}

func (repository *MemoryUser) FindByID(ctx context.Context, userID string) (*entity.User, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("http connection was cancel")
	default:
	}

	user, exist := repository.users[userID]
	if !exist {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
