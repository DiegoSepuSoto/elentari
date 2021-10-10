package auth

func (u *authUseCase) ValidateToken(accessToken string) (bool, error) {
	return u.authRepository.ValidateToken(accessToken)
}
