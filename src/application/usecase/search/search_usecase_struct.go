package search

import "elentari/src/domain/repository"

type UseCase struct {
	serviceRepository repository.ServiceRepository
	categoryRepository repository.CategoryRepository
}

func NewSearchUseCase(serviceRepository repository.ServiceRepository, categoryRepository repository.CategoryRepository) *UseCase {
	return &UseCase{
		serviceRepository:  serviceRepository,
		categoryRepository: categoryRepository,
	}
}
