package usecase

import errors "github.com/brunosilv96/simple_finance_api/internal/finance/usecase"

type DeleteCategory struct {
	CategoryRepository CategoryRepository
}

func NewDeleteCategory(categoryRepository CategoryRepository) *DeleteCategory {
	return &DeleteCategory{
		CategoryRepository: categoryRepository,
	}
}

func (usecase *DeleteCategory) Execute(id string) error {
	if id == "" {
		return &errors.InputCannotBeNil{
			Input: "id",
		}
	}

	error := usecase.CategoryRepository.Delete(id)
	if error != nil {
		return errors.CategoryNotFound
	}

	return nil
}
