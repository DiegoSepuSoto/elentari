package home

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type homeHandler struct {
	homeUseCase usecase.HomeUseCase
}

func NewHomeHandler(e *echo.Echo, homeUseCase usecase.HomeUseCase, jwtMiddleware echo.MiddlewareFunc) *homeHandler {
	h := &homeHandler{homeUseCase: homeUseCase}
	const basePath = "/v1/home"
	homeGroup := e.Group(basePath)
	homeGroup.Use(jwtMiddleware)
	homeGroup.GET("", h.getHome)

	return h
}

func (h homeHandler) getHome(c echo.Context) error {
	homeContent, err := h.homeUseCase.GetHomePage()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, homeContent)
}