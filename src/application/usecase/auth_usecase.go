package usecase

import "elentari/src/domain/models"

type AuthUseCase interface {
	ValidateToken(accessToken string) (bool, error)
	RefreshToken(accessToken string) (string, error)
	Login(email, password string) (*models.Student, error)
}
