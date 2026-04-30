package category

import (
	"context"

	entity "github.com/brunosilv96/simple_finance_api/internal/finance/category/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *entity.Category) error
	FindAll(ctx context.Context) ([]entity.Category, error)
	FindByID(ctx context.Context, id string) (*entity.Category, error)
	Delete(ctx context.Context, id string) error
}
