package entity

import (
	contactEntity "elentari/src/infrastructure/graphql/repository/iluvatar/contact/entity"
	personEntity "elentari/src/infrastructure/graphql/repository/iluvatar/person/entity"
	sharedEntity "elentari/src/shared/entity"
)

type ServicesWithPostsResponse struct {
	Services []*ServiceWithPosts `json:"servicios"`
}

type ServiceWithPosts struct {
	ID           string                      `json:"id"`
	Abbreviation string                      `json:"Sigla"`
	Posts        []*PostWithSummaryAndImages `json:"posts"`
}

type ServiceInPost struct {
	ID   string `json:"id"`
	Name string `json:"Nombre"`
}

type PostWithSummaryAndImages struct {
	ID      string                      `json:"id"`
	Title   string                      `json:"Titulo"`
	Summary string                      `json:"Resumen"`
	Image   *sharedEntity.IluvatarImage `json:"Imagen"`
}

type DetailedServiceResponse struct {
	DetailedService *DetailedService `json:"servicio"`
}

type DetailedService struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"Nombre"`
	Description string                      `json:"Descripcion"`
	Logo        *sharedEntity.IluvatarImage `json:"Logo"`
	Persons     []*personEntity.Person      `json:"personas"`
	Contacts    []*contactEntity.Contact    `json:"contactos"`
}

type Service struct {
	ID   string                      `json:"id"`
	Name string                      `json:"Nombre"`
	Logo *sharedEntity.IluvatarImage `json:"Logo"`
}

type ServicesResponse struct {
	Services []*Service `json:"servicios"`
}
