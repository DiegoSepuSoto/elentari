package service

import "elentari/src/domain/models"

func (u *UseCase) GetServicePostsPage(serviceID string) (*models.ServicePostsPage, error) {
	servicePosts, err := u.serviceRepository.GetServiceWithPosts(serviceID)
	if err != nil {
		return nil, err
	}

	return servicePosts, nil
}
