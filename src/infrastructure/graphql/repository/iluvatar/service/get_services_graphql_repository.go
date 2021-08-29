package service

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"log"
	"os"
)

const getServices = `
	query {
	  servicios(sort: "Nombre:asc") {
		id
		Nombre
		Descripcion
		Logo {
		  url
		}
      }
	}
`

func (r *serviceIluvatarRepository) GetServices() ([]*models.Service, error) {
	var res entity.ServicesResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getServices)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		log.Printf("Error getting services from GraphQL repository: %s", err.Error())
		return nil, err
	}

	return mapEntityServicesToServicesModel(res.Services), nil
}

func mapEntityServicesToServicesModel(entityServices []*entity.Service) []*models.Service {
	modelServices := make([]*models.Service, 0)

	for _, service := range entityServices {
		modelServices = append(modelServices, &models.Service{
			ID:      service.ID,
			Name:    service.Name,
			LogoURL: os.Getenv("ILUVATAR_URL") + service.Logo.URL,
		})
	}

	return modelServices
}