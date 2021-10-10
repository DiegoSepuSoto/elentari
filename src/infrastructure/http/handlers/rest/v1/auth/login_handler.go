package auth

import (
	"github.com/labstack/echo"
	"net/http"
)

const forbiddenUserError = "can't login user, wrong http code from response: 403"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *authHandler) Login(c echo.Context) error {
	var loginRequest *LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	student, err := h.authUseCase.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		if err.Error() == forbiddenUserError {
			return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, student)
}
