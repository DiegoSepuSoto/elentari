package category

import "elentari/src/domain/models"

func (u UseCase) GetCategories() ([]*models.Category, error) {
	categories, err := u.categoryRepository.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
