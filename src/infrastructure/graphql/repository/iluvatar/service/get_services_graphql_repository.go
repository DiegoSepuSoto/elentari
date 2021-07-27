package service

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
)

const getServices = `
	query {
	  servicios {
		id
		Nombre
		Sigla
	  }
	}
`

func (r serviceIluvatarRepository) GetServices() ([]*models.Service, error) {
	var res entity.ServicesResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getServices)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapServicesEntityToServicesModel(res.Services), nil
}

func mapServicesEntityToServicesModel(entityServices []*entity.Service) []*models.Service {
	domainServices := make([]*models.Service, 0)

	for _, service := range entityServices {
		domainServices = append(domainServices, &models.Service{
			ID:           service.ID,
			Name:         service.Name,
			Abbreviation: service.Abbreviation,
		})
	}

	return domainServices
}
