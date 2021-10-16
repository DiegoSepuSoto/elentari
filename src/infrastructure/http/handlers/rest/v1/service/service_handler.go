package service

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type serviceHandler struct {
	serviceUseCase usecase.ServiceUseCase
}

func NewServiceHandler(e *echo.Echo, serviceUseCase usecase.ServiceUseCase, jwtMiddleware echo.MiddlewareFunc) *serviceHandler {
	h := &serviceHandler{serviceUseCase: serviceUseCase}
	const basePath = "/v1/service"
	serviceGroup := e.Group(basePath)
	serviceGroup.Use(jwtMiddleware)
	serviceGroup.GET("/:id", h.getServicePage)
	serviceGroup.GET("/:id/posts", h.getServicePostsPage)

	return h
}

func (h serviceHandler) getServicePage(c echo.Context) error {
	serviceID := c.Param("id")
	servicePage, err := h.serviceUseCase.GetServicePage(serviceID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, servicePage)
}

func (h serviceHandler) getServicePostsPage(c echo.Context) error {
	serviceID := c.Param("id")
	servicePostsPage, err := h.serviceUseCase.GetServicePostsPage(serviceID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, servicePostsPage)
}