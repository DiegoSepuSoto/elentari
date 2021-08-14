package service

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"os"
)

const getServicesWithPosts = `
	query {
	  servicios {
		id
		Sigla
		posts {
		  id
		  Imagen {
			url
		  }
		  Resumen
		  Titulo
		}
	  }
	}
`

func (r serviceIluvatarRepository) GetServicesWithPosts() (*models.HomePage, error) {
	var res entity.ServicesWithPostsResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getServicesWithPosts)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapServicesWithPostEntityToHomeModel(res.Services), nil
}

func mapServicesWithPostEntityToHomeModel(entityServices []*entity.ServiceWithPosts) *models.HomePage {
	homePageServices := make([]*models.ServiceWithPosts, 0)

	for _, service := range entityServices {
		if len(service.Posts) > 0 {
			homePagePosts := make([]*models.PostForPreview, 0)
			for _, post := range service.Posts {
				homePagePosts = append(homePagePosts, &models.PostForPreview{
					ID:       post.ID,
					Title:    post.Title,
					Summary:  post.Summary,
					ImageURL: os.Getenv("ILUVATAR_URL") + post.Image.URL,
				})
			}
			homePageServices = append(homePageServices, &models.ServiceWithPosts{
				ID:              service.ID,
				Abbreviation:    service.Abbreviation,
				PostsForPreview: homePagePosts,
			})
		}
	}

	return &models.HomePage{ServicesWithPosts: homePageServices}
}