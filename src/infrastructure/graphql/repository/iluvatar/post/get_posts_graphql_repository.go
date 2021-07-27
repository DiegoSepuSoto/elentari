package post

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/post/entity"
	entity2 "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
)

const getPostsQuery = `
	query {
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
`

func (r postIluvatarRepository) GetPosts() ([]*models.Post, error) {
	var res entity.PostsResponse
	var errorResp entity2.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getPostsQuery)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapPostsEntityToPostsModel(res.Posts), nil
}

func mapPostsEntityToPostsModel(entityPosts []entity.Post) []*models.Post {
	domainPosts := make([]*models.Post, 0)

	for _, post := range entityPosts {
		domainPosts = append(domainPosts, &models.Post{
			ID:        post.ID,
			Title:     post.Title,
			Summary:   post.Summary,
			Body:      post.Body,
			ServiceID: post.Service.ID,
		})
	}

	return domainPosts
}
