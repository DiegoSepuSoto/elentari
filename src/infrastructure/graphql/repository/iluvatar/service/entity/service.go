package entity

import "elentari/src/infrastructure/graphql/repository/iluvatar/post/entity"

type ServicesWithPostsResponse struct {
	Services []*ServiceWithPosts `json:"servicios"`
}

type ServiceWithPosts struct {
	ID           string                             `json:"id"`
	Abbreviation string                             `json:"Sigla"`
	Posts        []*entity.PostWithSummaryAndImages `json:"posts"`
}
