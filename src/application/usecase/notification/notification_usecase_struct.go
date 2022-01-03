package notification

import "elentari/src/domain/repository"

type notificationUseCase struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationUseCase(notificationRepository repository.NotificationRepository) *notificationUseCase {
	return &notificationUseCase{notificationRepository: notificationRepository}
}
