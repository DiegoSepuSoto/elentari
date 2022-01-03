package repository

import "elentari/src/domain/models"

type NotificationRepository interface {
	GetNotifications() (*models.NotificationsPage, error)
}
