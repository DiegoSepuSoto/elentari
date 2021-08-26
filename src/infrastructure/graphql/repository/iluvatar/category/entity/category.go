package entity

import "elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"

type CategoriesResponse struct {
	Categories []*Category `json:"categorias"`
}

type CategoryWithPostsResponse struct {
	CategoryWithPosts *CategoryWithPosts `json:"categoria"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"Titulo"`
}

type CategoryWithPosts struct {
	Category
	Posts []*entity.PostWithSummaryAndImages `json:"posts"`
}
