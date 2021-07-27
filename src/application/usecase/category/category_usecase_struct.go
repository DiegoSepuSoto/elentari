package category

import "elentari/src/domain/repository"

type UseCase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepository repository.CategoryRepository) *UseCase {
	return &UseCase{categoryRepository: categoryRepository}
}