package post

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/post/entity"
	serviceEntity "elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"log"
	"os"
)

const getPostsByTerm = `
	query($term: String!) {
	  posts(where: {Cuerpo_contains: $term}) {
		id
		Titulo
		Resumen
		Imagen {
		  url
		}
	  }
	}
`

func (r *postIluvatarRepository) GetPostsByTerm(searchTerm string) (*models.PostSearchResult, error) {
	var res entity.PostsByTermResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getPostsByTerm)
	graphqlRequest.Var("term", searchTerm)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		log.Printf("Error getting posts by term: %s from GraphQL repository: %s", searchTerm, err.Error())
		return nil, err
	}

	return &models.PostSearchResult{Posts: mapEntityPostsByTermToModelPostsByTerm(res.Posts)}, nil
}

func mapEntityPostsByTermToModelPostsByTerm(entityPosts []*serviceEntity.PostWithSummaryAndImages) []*models.PostForPreview {
	modelPosts := make([]*models.PostForPreview, 0)

	for _, post := range entityPosts {
		modelPosts = append(modelPosts, &models.PostForPreview{
			ID:       post.ID,
			Title:    post.Title,
			Summary:  post.Summary,
			ImageURL: os.Getenv("ILUVATAR_URL") + post.Image.URL,
		})
	}

	return modelPosts
}
