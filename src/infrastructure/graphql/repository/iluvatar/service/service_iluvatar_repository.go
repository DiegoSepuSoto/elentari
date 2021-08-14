package service

import (
	"fmt"
	"github.com/pzentenoe/graphql-client"
	"os"
)

type serviceIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewServiceIluvatarRepository(graphqlClient *graphql.Client) *serviceIluvatarRepository {
	return &serviceIluvatarRepository{graphqlClient: graphqlClient}
}

func (r *serviceIluvatarRepository) fillGraphqlHeaders(graphqlRequest *graphql.GraphqlRequest) {
	iluvatarToken := os.Getenv("ILUVATAR_TOKEN")
	graphqlRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", iluvatarToken))
}