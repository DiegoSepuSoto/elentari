package usecase

import "elentari/src/domain/models"

type PostUseCase interface {
	 GetPosts() ([]*models.Post, error)
	 GetPostsByServiceID(serviceID string) ([]*models.Post, error)
}