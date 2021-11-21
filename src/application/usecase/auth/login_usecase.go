package auth

import (
	"elentari/src/domain/models"
	"elentari/src/domain/models/requests"
)

func (u *authUseCase) Login(loginRequest *requests.LoginRequest) (*models.Student, error) {
	return u.authRepository.Login(loginRequest)
}