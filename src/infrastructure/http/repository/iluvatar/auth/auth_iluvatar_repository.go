package auth

import (
	client "github.com/pzentenoe/httpclient-call-go"
)

type authIluvatarRepository struct {
	httpClientCall *client.HTTPClientCall
}

func NewAuthIluvatarRepository(httpClientCall *client.HTTPClientCall) *authIluvatarRepository {
	return &authIluvatarRepository{httpClientCall: httpClientCall}
}