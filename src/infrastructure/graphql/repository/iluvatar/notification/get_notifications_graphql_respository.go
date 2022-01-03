package notification

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/notification/entity"
	errorEntity "elentari/src/shared/entity"
	"fmt"
	"github.com/pzentenoe/graphql-client"
	"log"
	"os"
)

const getNotifications = `
	query {
	  avisos(sort: "createdAt:desc", limit: 10) {
		id,
		Titulo,
		Cuerpo,
		servicio {
			id,
			Nombre,
			Logo {
				url
			}
		}
	  }
	}
`

func (r *notificationIluvatarRepository) GetNotifications() (*models.NotificationsPage, error) {
	var res entity.NotificationsResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getNotifications)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		log.Printf("Error getting notifications")
		return nil, err
	}

	return &models.NotificationsPage{Notifications: mapNotificationsResponseToDomain(res.Notifications)}, nil
}

func mapNotificationsResponseToDomain(notifications []*entity.Notification) []*models.Notification {
	domainNotifications := make([]*models.Notification, 0)

	for _, notification := range notifications {
		domainNotifications = append(domainNotifications, &models.Notification{
			ID:              notification.ID,
			Title:           notification.Title,
			Body:            notification.Body,
			ServiceName:     notification.Service.Name,
			ServiceImageURL: fmt.Sprintf("%s%s", os.Getenv("ILUVATAR_CMS_HOST"), notification.Service.Logo.URL),
		})
	}

	return domainNotifications
}
