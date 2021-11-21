package auth

import (
	"elentari/src/domain/models/requests"
	"github.com/labstack/echo"
	"net/http"
)

const forbiddenUserError = "can't login user, wrong http code from response: 403"

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
