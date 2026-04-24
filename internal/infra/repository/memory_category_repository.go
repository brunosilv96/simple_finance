package repository

import (
	"errors"

	category "github.com/brunosilv96/simple_finance_api/internal/finance/category/entity"
)

type MemoryCategory struct {
	categories map[string]category.Category
}

// MemoryCategory Constructor
func NewMemoryCategory() *MemoryCategory {
	return &MemoryCategory{
		categories: make(map[string]category.Category),
	}
}

func (repository *MemoryCategory) Save(category *category.Category) error {
	repository.categories[category.ID] = *category

	return nil
}

func (repository *MemoryCategory) FindAll() ([]category.Category, error) {
	// Cria um slice do tamanho do map
	categories := make([]category.Category, 0, len(repository.categories))

	for _, category := range repository.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *MemoryCategory) FindByID(id string) (*category.Category, error) {
	category, exists := repository.categories[id]
	if !exists {
		return nil, errors.New("category not found")
	}

	return &category, nil
}

func (repository *MemoryCategory) Delete(id string) error {
	category, exists := repository.categories[id]
	if !exists {
		return errors.New("category not found")
	}

	delete(repository.categories, category.ID)

	return nil
}
