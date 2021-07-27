package post

import "elentari/src/domain/repository"

type UseCase struct {
	postRepository repository.PostRepository
}

func NewPostUseCase(postRepository repository.PostRepository) *UseCase {
	return &UseCase{postRepository: postRepository}
}