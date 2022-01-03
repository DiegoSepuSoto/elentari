package notification

import (
	"fmt"
	"github.com/pzentenoe/graphql-client"
	"os"
)

type notificationIluvatarRepository struct {
	graphqlClient *graphql.Client
}

func NewNotificationIluvatarRepository(graphqlClient *graphql.Client) *notificationIluvatarRepository {
	return &notificationIluvatarRepository{graphqlClient: graphqlClient}
}

func (r *notificationIluvatarRepository) fillGraphqlHeaders(graphqlRequest *graphql.GraphqlRequest) {
	iluvatarToken := os.Getenv("ILUVATAR_TOKEN")
	graphqlRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", iluvatarToken))
}
