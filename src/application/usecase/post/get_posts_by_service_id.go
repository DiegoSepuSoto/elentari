package post

import "elentari/src/domain/models"

func (u UseCase) GetPostsByServiceID(serviceID string) ([]*models.Post, error) {
	posts, err := u.postRepository.GetPostsByServiceID(serviceID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
