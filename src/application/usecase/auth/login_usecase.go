package auth

import "elentari/src/domain/models"

func (u *authUseCase) Login(email, password string) (*models.Student, error) {
	return u.authRepository.Login(email, password)
}