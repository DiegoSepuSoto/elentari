package catalog

import "elentari/src/domain/models"

func (u *UseCase) GetCatalogPage() (*models.CatalogPage, error) {
	services, err := u.serviceRepository.GetServices()
	if err != nil {
		return nil, err
	}

	categories, err := u.categoryRepository.GetCategories()
	if err != nil {
		return nil, err
	}

	return &models.CatalogPage{
		Services:   services,
		Categories: categories,
	}, nil
}