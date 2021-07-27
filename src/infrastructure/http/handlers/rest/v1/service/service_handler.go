package service

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type serviceHandler struct {
	serviceUseCase usecase.ServiceUseCase
}

func NewServiceHandler(e *echo.Echo, serviceUseCase usecase.ServiceUseCase) *serviceHandler {
	h := &serviceHandler{serviceUseCase: serviceUseCase}
	const basePath = "/v1/services"
	postGroup := e.Group(basePath)
	postGroup.GET("", h.getServices)

	return h
}

func (h serviceHandler) getServices(c echo.Context) error {
	services, err := h.serviceUseCase.GetServices()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, services)
}