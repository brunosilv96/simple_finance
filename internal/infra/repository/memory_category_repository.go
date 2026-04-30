package repository

import (
	"context"
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

func (repository *MemoryCategory) Save(ctx context.Context, category *category.Category) error {
	select {
	case <-ctx.Done():
		return errors.New("http connection was cancel")
	default:
	}

	repository.categories[category.ID] = *category

	return nil
}

func (repository *MemoryCategory) FindAll(ctx context.Context) ([]category.Category, error) {
	select {
	case <-ctx.Done():
		return []category.Category{}, errors.New("http connection was cancel")
	default:
	}

	// Cria um slice do tamanho do map
	categories := make([]category.Category, 0, len(repository.categories))

	for _, category := range repository.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *MemoryCategory) FindByID(ctx context.Context, id string) (*category.Category, error) {
	select {
	case <-ctx.Done():
		return &category.Category{}, errors.New("http connection was cancel")
	default:
	}

	category, exists := repository.categories[id]
	if !exists {
		return nil, errors.New("category not found")
	}

	return &category, nil
}

func (repository *MemoryCategory) Delete(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return errors.New("http connection was cancel")
	default:
	}

	category, exists := repository.categories[id]
	if !exists {
		return errors.New("category not found")
	}

	delete(repository.categories, category.ID)

	return nil
}
