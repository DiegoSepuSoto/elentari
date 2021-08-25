package entity

import (
	categoryEntity "elentari/src/infrastructure/graphql/repository/iluvatar/category/entity"
	serviceEntity "elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	sharedEntity "elentari/src/shared/entity"
)

type DetailedPostResponse struct {
	DetailedPost *DetailedPost `json:"post"`
}

type DetailedPost struct {
	ID         string                      `json:"id"`
	Title      string                      `json:"Titulo"`
	Image      *sharedEntity.IluvatarImage `json:"Imagen"`
	Body       string                       `json:"Cuerpo"`
	Service    *serviceEntity.ServiceInPost `json:"servicio"`
	Categories []*categoryEntity.Category   `json:"categorias"`
}
