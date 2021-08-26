package category

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/category/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"os"
)

const getCategoryPosts = `
	query($id: ID!) {
	  categoria(id: $id) {
		id
		Titulo
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

func (r *categoryIluvatarRepository) GetCategoryWithPosts(categoryID string) (*models.CategoryPostsPage, error) {
	var res entity.CategoryWithPostsResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getCategoryPosts)
	graphqlRequest.Var("id", categoryID)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapEntityCategoryWithPostsToModelCategoryWithPosts(res.CategoryWithPosts), nil
}

func mapEntityCategoryWithPostsToModelCategoryWithPosts(entityCategoryWithPosts *entity.CategoryWithPosts) *models.CategoryPostsPage {
	posts := make([]*models.PostForPreview, 0)

	for _, post := range entityCategoryWithPosts.Posts {
		posts = append(posts, &models.PostForPreview{
			ID:       post.ID,
			Title:    post.Title,
			Summary:  post.Summary,
			ImageURL: os.Getenv("ILUVATAR_URL") + post.Image.URL,
		})
	}

	return &models.CategoryPostsPage{
		ID:              entityCategoryWithPosts.ID,
		Name:            entityCategoryWithPosts.Name,
		PostsForPreview: posts,
	}
}