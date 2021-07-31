package entity

import "elentari/src/shared/entity"

type PostWithSummaryAndImages struct {
	ID      string                `json:"id"`
	Title   string                `json:"Titulo"`
	Summary string                `json:"Resumen"`
	Image   *entity.IluvatarImage `json:"Imagen"`
}
