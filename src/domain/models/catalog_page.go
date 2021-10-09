package models

type Service struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CatalogPage struct {
	Services []*Service `json:"services"`
	Categories []*Category `json:"categories"`
}
