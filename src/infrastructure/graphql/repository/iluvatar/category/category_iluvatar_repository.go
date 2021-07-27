package category

import "github.com/pzentenoe/graphql-client"

type categoryIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewCategoryIluvatarRepository(graphqlClient *graphql.Client) *categoryIluvatarRepository {
	return &categoryIluvatarRepository{graphqlClient: graphqlClient}
}
