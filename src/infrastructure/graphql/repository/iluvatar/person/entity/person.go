package entity

import "elentari/src/shared/entity"

type Person struct {
	ID             string                `json:"id"`
	Name           string                `json:"Nombre"`
	Charge         string                `json:"Cargo"`
	Email          string                `json:"Correo"`
	ProfilePicture *entity.IluvatarImage `json:"Imagen"`
}
