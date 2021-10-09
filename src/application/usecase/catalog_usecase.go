package usecase

import "elentari/src/domain/models"

type CatalogUseCase interface {
	GetCatalogPage() (*models.CatalogPage, error)
}
