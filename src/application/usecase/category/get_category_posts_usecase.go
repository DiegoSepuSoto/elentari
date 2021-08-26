package category

import "elentari/src/domain/models"

func (u *UseCase) GetCategoryPostsPage(categoryID string) (*models.CategoryPostsPage, error) {
	categoryWithPosts, err := u.categoryRepository.GetCategoryWithPosts(categoryID)
	if err != nil {
		return nil, err
	}

	return categoryWithPosts, nil
}