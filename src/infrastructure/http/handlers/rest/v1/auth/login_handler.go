package auth

import (
	"elentari/src/domain/models/requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

const forbiddenUserError = "can't login user, wrong http code from response: 403"

// @Summary Inicio de Sesión
// @Description Permite a los estudiantes iniciar sesión con sus credenciales de Pasaporte.UTEM
// @Tags BFF V1 - Autenticación
// @Accept json
// @Produce json
// @Success 200 {object} models.Student "OK"
// @Failure 400 {object} map[string]interface{} "BadRequest"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/auth/login [post]
func (h *authHandler) Login(c echo.Context) error {
	var loginRequest *requests.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	student, err := h.authUseCase.Login(loginRequest)
	if err != nil {
		if err.Error() == forbiddenUserError {
			return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}
