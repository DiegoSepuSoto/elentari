package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Actualización de Token
// @Description Para reducir la cantidad de veces que un usuario debe iniciar sesión, se diponibiliza un endpoint para actualizar el token de consultas a los diferentes servicios
// @Tags BFF V1 - Autenticación
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "OK"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/auth/refresh-token [post]
func (h *authHandler) RefreshToken(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)

	newToken, err := h.authUseCase.RefreshToken(token)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": fmt.Sprintf("there was an error getting new token [%s]", err.Error())})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": newToken})
}
