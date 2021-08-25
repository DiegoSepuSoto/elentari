package search

import "elentari/src/domain/models"

func (u *UseCase) GetSearchPage() (*models.SearchPage, error) {
	services, err := u.serviceRepository.GetServices()
	if err != nil {
		return nil, err
	}

	categories, err := u.categoryRepository.GetCategories()
	if err != nil {
		return nil, err
	}

	return &models.SearchPage{
		Services:   services,
		Categories: categories,
	}, nil
}