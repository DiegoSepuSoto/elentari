package post

import (
	"fmt"
	"github.com/pzentenoe/graphql-client"
	"os"
)

type postIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewPostIluvatarRepository(graphqlClient *graphql.Client) *postIluvatarRepository {
	return &postIluvatarRepository{graphqlClient: graphqlClient}
}


func (r *postIluvatarRepository) fillGraphqlHeaders(graphqlRequest *graphql.GraphqlRequest) {
	iluvatarToken := os.Getenv("ILUVATAR_TOKEN")
	graphqlRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", iluvatarToken))
}