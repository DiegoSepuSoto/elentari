package catalog

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type catalogHandler struct {
	catalogUseCase usecase.CatalogUseCase
}

func NewCatalogHandler(e *echo.Echo, catalogUseCase usecase.CatalogUseCase, jwtMiddleware echo.MiddlewareFunc) *catalogHandler {
	h := &catalogHandler{catalogUseCase: catalogUseCase}
	const basePath = "/v1/catalog"
	catalogGroup := e.Group(basePath)
	catalogGroup.Use(jwtMiddleware)
	catalogGroup.GET("", h.getCatalog)

	return h
}

// @Summary Pantalla de Catálogo
// @Description Devuelve el contenido a mostrar en la pantalla de catélogo
// @Tags BFF V1 - Catálogo
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Success 200 {object} models.CatalogPage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/catalog [get]
func (h *catalogHandler) getCatalog(c echo.Context) error {
	catalogPage, err := h.catalogUseCase.GetCatalogPage()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, catalogPage)
}