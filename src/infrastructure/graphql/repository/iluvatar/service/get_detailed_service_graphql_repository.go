package service

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
)

const getDetailedService = `
	query($id: ID!) {
	  servicio(id: $id) {
		id
		Nombre
		Logo {
		  url
		}
		Descripcion
	  }
	}
`
func (r serviceIluvatarRepository) GetDetailedService(serviceID string) (*models.ServicePage, error) {
	var res entity.DetailedServiceResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getDetailedService)
	graphqlRequest.Var("id", serviceID)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapEntityDetailedServiceToServicePageModel(res.DetailedService), nil
}

func mapEntityDetailedServiceToServicePageModel(service *entity.DetailedService) *models.ServicePage {
	return &models.ServicePage{
		ID:          service.ID,
		LogoURL:     service.Logo.URL,
		Name:        service.Name,
		Description: service.Description,
	}
}
