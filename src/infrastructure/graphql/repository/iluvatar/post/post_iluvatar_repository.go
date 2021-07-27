package post

import "github.com/pzentenoe/graphql-client"

type postIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewPostIluvatarRepository(graphqlClient *graphql.Client) *postIluvatarRepository {
	return &postIluvatarRepository{graphqlClient: graphqlClient}
}
