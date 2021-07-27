package category

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/category/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
)

const getCategoriesQuery = `
	query {
	  categorias {
		id
		Nombre
	  }
	}
`

func (r categoryIluvatarRepository) GetCategories() ([]*models.Category, error) {
	var res entity.CategoriesResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getCategoriesQuery)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapCategoriesEntityToCategoriesModel(res.Categories), nil
}

func mapCategoriesEntityToCategoriesModel(entityCategories []*entity.Category) []*models.Category {
	domainCategories := make([]*models.Category, 0)
	for _, category := range entityCategories {
		domainCategories = append(domainCategories, &models.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return domainCategories
}
