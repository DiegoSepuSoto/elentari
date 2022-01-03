package usecase

import "elentari/src/domain/models"

type NotificationUseCase interface {
	GetNotifications() (*models.NotificationsPage, error)
}
