package repository

import "elentari/src/domain/models"

type ServiceRepository interface {
	GetServicesWithPosts() (*models.HomePage, error)
}
