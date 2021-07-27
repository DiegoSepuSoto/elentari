package service

import "elentari/src/domain/models"

func (u UseCase) GetServices() ([]*models.Service, error) {
	services, err := u.serviceRepository.GetServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}
