package service

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"os"
)

const getServiceWithPosts = `
	query($id: ID!) {
	  servicio(id: $id) {
		id
		Nombre
		posts(sort: "createdAt:desc") {
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

func (r serviceIluvatarRepository) GetServiceWithPosts(serviceID string) (*models.ServicePostsPage, error) {
	var res entity.ServiceWithPostsResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getServiceWithPosts)
	graphqlRequest.Var("id", serviceID)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapEntityServiceWithPostsToModelServiceWithPosts(res.Service), nil
}

func mapEntityServiceWithPostsToModelServiceWithPosts(serviceWithPostsEntity *entity.ServiceWithPosts) *models.ServicePostsPage {
	posts := make([]*models.PostForPreview, 0)

	for _, post := range serviceWithPostsEntity.Posts {
		posts = append(posts, &models.PostForPreview{
			ID:       post.ID,
			Title:    post.Title,
			Summary:  post.Summary,
			ImageURL: os.Getenv("ILUVATAR_URL") + post.Image.URL,
		})
	}

	return &models.ServicePostsPage{
		ID:              serviceWithPostsEntity.ID,
		Name:            serviceWithPostsEntity.Name,
		PostsForPreview: posts,
	}
}

