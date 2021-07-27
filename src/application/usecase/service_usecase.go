package usecase

import "elentari/src/domain/models"

type ServiceUseCase interface {
	GetServices() ([]*models.Service, error)
}
