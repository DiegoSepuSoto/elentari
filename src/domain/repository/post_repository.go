package repository

import "elentari/src/domain/models"

type PostRepository interface {
	GetPosts() ([]*models.Post, error)
	GetPostsByServiceID(serviceID string) ([]*models.Post, error)
}
