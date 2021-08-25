package repository

import "elentari/src/domain/models"

type PostRepository interface {
	GetDetailedPost(postID string) (*models.PostPage, error)
	GetPostsByTerm(searchTerm string) (*models.PostSearchResult, error)
}
