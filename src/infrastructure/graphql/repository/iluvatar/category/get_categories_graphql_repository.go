package category

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/category/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
)

const getCategories = `
	query {
	  categorias(sort: "Nombre:asc") {
		id
		Titulo
	  }
	}
`

func (r *categoryIluvatarRepository) GetCategories() ([]*models.Category, error) {
	var res entity.CategoriesResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getCategories)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapEntityCategoriesToModelCategories(res.Categories), nil
}

func mapEntityCategoriesToModelCategories(entityCategories []*entity.Category) []*models.Category {
	modelsCategory := make([]*models.Category, 0)

	for _, category := range entityCategories {
		modelsCategory = append(modelsCategory, &models.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return modelsCategory
}