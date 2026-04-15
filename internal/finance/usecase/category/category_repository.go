package usecase

import "github.com/brunosilv96/simple_finance_api/internal/finance/entity"

type CategoryRepository interface {
	Save(category *entity.Category) error
	FindAll() ([]entity.Category, error)
	FindByID(id string) (*entity.Category, error)
	Delete(id string) error
}
