package repository

import "elentari/src/domain/models"

type CategoryRepository interface {
	GetCategories() ([]*models.Category, error)
	GetCategoryWithPosts(categoryID string) (*models.CategoryPostsPage, error)
}
