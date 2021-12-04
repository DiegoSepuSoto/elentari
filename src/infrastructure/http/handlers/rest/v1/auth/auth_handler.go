package auth

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
)

const basePath = "/v1/auth"

type authHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(e *echo.Echo, authUseCase usecase.AuthUseCase) *authHandler {
	h := &authHandler{authUseCase: authUseCase}

	authHandlerWithoutAuthentication := e.Group(basePath)
	authHandlerWithoutAuthentication.POST("/login", h.Login)
	authHandlerWithoutAuthentication.POST("/refresh-token", h.RefreshToken)
	authHandlerWithoutAuthentication.POST("/validate-token", h.validateToken)

	return h
}