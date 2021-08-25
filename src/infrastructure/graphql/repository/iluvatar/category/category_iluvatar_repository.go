package category

import (
	"fmt"
	"github.com/pzentenoe/graphql-client"
	"os"
)

type categoryIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewCategoryIluvatarRepository(graphqlClient *graphql.Client) *categoryIluvatarRepository {
	return &categoryIluvatarRepository{graphqlClient: graphqlClient}
}

func (r *categoryIluvatarRepository) fillGraphqlHeaders(graphqlRequest *graphql.GraphqlRequest) {
	iluvatarToken := os.Getenv("ILUVATAR_TOKEN")
	graphqlRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", iluvatarToken))
}