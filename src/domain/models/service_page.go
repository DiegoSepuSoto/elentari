package models

type ServicePage struct {
	ID string `json:"id"`
	LogoURL string `json:"logo_url"`
	Name string `json:"name"`
	Description string `json:"description"`
}