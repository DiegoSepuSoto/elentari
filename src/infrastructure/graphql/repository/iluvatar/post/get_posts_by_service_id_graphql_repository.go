package post

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/post/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
)

const getPostsByServiceIDQuery = `
	query($id: ID!) {
	  servicio(id: $id) {
		id,
		posts {
			id
			Titulo,
			Resumen,
			Cuerpo,
			servicio {
			  id
			}
		  }
	  }
	}
`

func (r postIluvatarRepository) GetPostsByServiceID(serviceID string) ([]*models.Post, error) {
	var res entity.PostsByServiceIDResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getPostsByServiceIDQuery)
	graphqlRequest.Var("id", serviceID)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapPostsEntityToPostsModel(res.Service.Posts), nil
}
