package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Validación de Token
// @Description Con el fin de detectar la validez de un token se disponibiliza este endpoint para verificar su integridad
// @Tags BFF V1 - Autenticación
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /v1/auth/validate-token [post]
func (h *authHandler) validateToken(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	isValid, err := h.authUseCase.ValidateToken(token)

	if err != nil || !isValid {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "valid token", "valid": true})
}
