package usecase

import "elentari/src/domain/models"

type CategoryUseCase interface {
	GetCategoryPostsPage(categoryID string) (*models.CategoryPostsPage, error)
}