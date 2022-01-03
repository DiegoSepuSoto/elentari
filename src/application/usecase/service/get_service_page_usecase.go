package service

import "elentari/src/domain/models"

func (u *UseCase) GetServicePage(serviceID string) (*models.ServicePage, error) {
	return u.serviceRepository.GetDetailedService(serviceID)
}
