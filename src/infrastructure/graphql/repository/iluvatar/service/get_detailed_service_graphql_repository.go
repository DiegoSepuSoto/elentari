package service

import (
	"context"
	"elentari/src/domain/models"
	"elentari/src/infrastructure/graphql/repository/iluvatar/service/entity"
	errorEntity "elentari/src/shared/entity"
	"github.com/pzentenoe/graphql-client"
	"os"
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
		personas {
		  id
		  Nombre
		  Cargo
		  Correo
		  Imagen {
			url
		  }
		}
		contactos {
		  id
		  Nombre
		  Tipo
		  Link
		}
	  }
	}
`
func (r serviceIluvatarRepository) GetDetailedService(serviceID string) (*models.ServicePage, error) {
	var res entity.DetailedServiceResponse
	var errorResp errorEntity.ErrorResponse
	graphqlRequest := graphql.NewGraphqlRequest(getDetailedService)
	graphqlRequest.Var("id", serviceID)

	r.fillGraphqlHeaders(graphqlRequest)

	err := r.graphqlClient.Run(context.Background(), graphqlRequest, &res, &errorResp)
	if err != nil {
		return nil, err
	}

	return mapEntityDetailedServiceToServicePageModel(res.DetailedService), nil
}

func mapEntityDetailedServiceToServicePageModel(service *entity.DetailedService) *models.ServicePage {
	var profilePictureLink string
	modelPersons := make([]*models.Person, 0)
	modelContacts := make([]*models.Contact, 0)

	for _, person := range service.Persons {
		if person.ProfilePicture == nil {
			profilePictureLink = ""
		} else {
			profilePictureLink = os.Getenv("ILUVATAR_URL") + person.ProfilePicture.URL
		}
		modelPersons = append(modelPersons, &models.Person{
			ID:             person.ID,
			Name:           person.Name,
			Charge:         person.Charge,
			Email:          person.Email,
			ProfilePicture: profilePictureLink,
		})
	}

	for _, contact := range service.Contacts {
		modelContacts = append(modelContacts, &models.Contact{
			ID:   contact.ID,
			Name: contact.Name,
			Type: contact.Type,
			Link: contact.Link,
		})
	}

	return &models.ServicePage{
		ID:          service.ID,
		LogoURL:     os.Getenv("ILUVATAR_URL") + service.Logo.URL,
		Name:        service.Name,
		Description: service.Description,
		Persons:     modelPersons,
		Contacts:    modelContacts,
	}
}
