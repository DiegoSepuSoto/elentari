package usecase

import "elentari/src/domain/models"

type CategoryUseCase interface {
	GetCategories() ([]*models.Category, error)
}