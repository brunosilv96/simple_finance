package usecase

import "github.com/brunosilv96/simple_finance_api/internal/finance/entity"

type FindAllCategories struct {
	CategoryRepository CategoryRepository
}

func NewListCategories(categoryRepository CategoryRepository) *FindAllCategories {
	return &FindAllCategories{
		CategoryRepository: categoryRepository,
	}
}

func (usecase *FindAllCategories) Execute() ([]entity.Category, error) {
	categories, err := usecase.CategoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
