package repository

import "elentari/src/domain/models"

type ServiceRepository interface {
	GetServices() ([]*models.Service, error)
}
