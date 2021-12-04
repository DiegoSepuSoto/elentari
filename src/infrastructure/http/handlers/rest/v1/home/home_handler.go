package home

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
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

// @Summary Pantalla de Inicio
// @Description Devuelve el contenido a mostrar en la pantalla de inicio de la aplicación
// @Tags BFF V1 - Pantalla Principal
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Success 200 {object} models.HomePage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/home [get]
func (h homeHandler) getHome(c echo.Context) error {
	homeContent, err := h.homeUseCase.GetHomePage()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, homeContent)
}