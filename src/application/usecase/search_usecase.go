package usecase

import "elentari/src/domain/models"

type SearchUseCase interface {
	GetSearchPage() (*models.SearchPage, error)
}
