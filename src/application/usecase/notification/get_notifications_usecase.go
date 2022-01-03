package notification

import "elentari/src/domain/models"

func (u *notificationUseCase) GetNotifications() (*models.NotificationsPage, error) {
	return u.notificationRepository.GetNotifications()
}
