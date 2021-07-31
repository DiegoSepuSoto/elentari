package post

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/post/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"os"
)

const getDetailedPost = `
	query($id: ID!) {
	  post(id: $id) {
		id
		Imagen {
		  url
		}
		Titulo
		categorias {
		  id
		  Nombre
		}
		Cuerpo
		servicio {
		  id
		}
	  }
	}
`

func (r postIluvatarRepository) GetDetailedPost(postID string) (*models.PostPage, error) {
	var res entity.DetailedPostResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getDetailedPost)
	graphqlRequest.Var("id", postID)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapEntityDetailedPostToPostPageModel(res.DetailedPost), nil
}

func mapEntityDetailedPostToPostPageModel(entityPost *entity.DetailedPost) *models.PostPage {
	postCategories := make([]*models.PostCategory, 0)

	for _, category := range entityPost.Categories {
		postCategories = append(postCategories, &models.PostCategory{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return &models.PostPage{
		ID:         entityPost.ID,
		Title:      entityPost.Title,
		Image:      os.Getenv("ILUVATAR_URL") + entityPost.Image.URL,
		Body:       entityPost.Body,
		ServiceID:  entityPost.Service.ID,
		Categories: postCategories,
	}
}