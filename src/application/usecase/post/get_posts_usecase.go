package post

import "elentari/src/domain/models"

func (u *UseCase) GetPosts() ([]*models.Post, error) {
	posts, err := u.postRepository.GetPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}