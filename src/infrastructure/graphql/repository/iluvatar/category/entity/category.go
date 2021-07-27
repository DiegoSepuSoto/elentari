package entity

type CategoriesResponse struct {
	Categories []*Category `json:"categorias"`
}

type Category struct {
	ID string `json:"id"`
	Name string `json:"Nombre"`
}
