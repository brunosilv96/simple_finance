package user

import (
	"context"

	entity "github.com/brunosilv96/simple_finance_api/internal/finance/user/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, userID string) (*entity.User, error)
}
