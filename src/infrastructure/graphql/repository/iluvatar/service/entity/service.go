package entity

type ServicesResponse struct {
	Services []*Service `json:"servicios"`
}

type Service struct {
	ID           string `json:"id"`
	Name         string `json:"Nombre"`
	Abbreviation string `json:"Sigla"`
}
