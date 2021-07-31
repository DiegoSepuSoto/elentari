package usecase

import "elentari/src/domain/models"

type HomeUseCase interface {
	GetHomePage() (*models.HomePage, error)
}