package home

import "elentari/src/domain/repository"

type UseCase struct {
	serviceRepository repository.ServiceRepository
}

func NewHomeUseCase(serviceRepository repository.ServiceRepository) *UseCase {
	return &UseCase{serviceRepository: serviceRepository}
}