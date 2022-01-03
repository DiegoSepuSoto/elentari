package entity

import serviceEntity "elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"

type NotificationsResponse struct {
	Notifications []*Notification `json:"avisos"`
}

type Notification struct {
	ID      string                               `json:"id"`
	Title   string                               `json:"Titulo"`
	Body    string                               `json:"Cuerpo"`
	Service *serviceEntity.ServiceInNotification `json:"servicio"`
}
