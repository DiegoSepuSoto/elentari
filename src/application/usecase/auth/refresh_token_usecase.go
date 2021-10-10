package auth

func (u *authUseCase) RefreshToken(accessToken string) (string, error) {
	return u.authRepository.RefreshToken(accessToken)
}