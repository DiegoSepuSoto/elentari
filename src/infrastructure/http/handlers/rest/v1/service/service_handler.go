package service

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
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

// @Summary Pantalla de Servicio
// @Description Devuelve el detalle de un servicio estudiantil para visualizarlo en la aplicación
// @Tags BFF V1 - Servicios Estudiantiles
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Param id path string true "123456"
// @Success 200 {object} models.ServicePage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/service/{id} [get]
func (h serviceHandler) getServicePage(c echo.Context) error {
	serviceID := c.Param("id")
	servicePage, err := h.serviceUseCase.GetServicePage(serviceID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, servicePage)
}

// @Summary Publicaciones de un Servicio
// @Description Devuelve todas las publicaciones creadas por un servicio estudiantil especificado
// @Tags BFF V1 - Servicios Estudiantiles
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Param id path string true "123456"
// @Success 200 {object} models.ServicePostsPage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/service/{id}/posts [get]
func (h serviceHandler) getServicePostsPage(c echo.Context) error {
	serviceID := c.Param("id")
	servicePostsPage, err := h.serviceUseCase.GetServicePostsPage(serviceID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, servicePostsPage)
}