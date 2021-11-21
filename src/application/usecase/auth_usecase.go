package usecase

import (
	"elentari/src/domain/models"
	"elentari/src/domain/models/requests"
)

type AuthUseCase interface {
	ValidateToken(accessToken string) (bool, error)
	RefreshToken(accessToken string) (string, error)
	Login(loginRequest *requests.LoginRequest) (*models.Student, error)
}
