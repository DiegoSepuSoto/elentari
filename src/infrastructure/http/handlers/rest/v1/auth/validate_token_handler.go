package auth

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *authHandler) validateToken(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	isValid, err := h.authUseCase.ValidateToken(token)

	if err != nil || !isValid {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "valid token", "valid": true})
}
