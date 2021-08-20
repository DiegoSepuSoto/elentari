package models

type Person struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Charge         string `json:"charge"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture_url"`
}

type Contact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Link string `json:"link"`
}

type ServicePage struct {
	ID          string     `json:"id"`
	LogoURL     string     `json:"logo_url"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Persons     []*Person  `json:"persons"`
	Contacts    []*Contact `json:"contacts"`
}
