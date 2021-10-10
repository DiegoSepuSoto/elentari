package auth

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (h *authHandler) RefreshToken(c echo.Context) error {
	token := c.Request().Header.Get(echo.HeaderAuthorization)

	newToken, err := h.authUseCase.RefreshToken(token)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": fmt.Sprintf("there was an error getting new token [%s]", err.Error())})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": newToken})
}
