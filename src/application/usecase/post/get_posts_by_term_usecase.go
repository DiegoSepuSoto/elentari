package post

import "elentari/src/domain/models"

func (u *UseCase) GetPostsByTerm(searchTerm string) (*models.PostSearchResult, error) {
	postsByTerm, err := u.postRepository.GetPostsByTerm(searchTerm)
	if err != nil {
		return nil, err
	}

	return postsByTerm, nil
}
