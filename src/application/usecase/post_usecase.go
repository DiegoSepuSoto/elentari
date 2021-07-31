package usecase

import "elentari/src/domain/models"

type PostUseCase interface {
	GetPostPage(postID string) (*models.PostPage, error)
}
