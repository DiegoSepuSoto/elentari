package repository

import (
	"elentari/src/domain/models"
	"elentari/src/domain/models/requests"
)

type AuthRepository interface {
	ValidateToken(accessToken string) (bool, error)
	RefreshToken(accessToken string) (string, error)
	Login(loginRequest *requests.LoginRequest) (*models.Student, error)
}
