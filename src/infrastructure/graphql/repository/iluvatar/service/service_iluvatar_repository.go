package service

import "github.com/pzentenoe/graphql-client"

type serviceIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewServiceIluvatarRepository(graphqlClient *graphql.Client) *serviceIluvatarRepository {
	return &serviceIluvatarRepository{graphqlClient: graphqlClient}
}