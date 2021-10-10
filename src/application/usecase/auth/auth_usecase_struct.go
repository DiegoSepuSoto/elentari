package auth

import "elentari/src/domain/repository"

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) *authUseCase {
	return &authUseCase{authRepository: authRepository}
}
