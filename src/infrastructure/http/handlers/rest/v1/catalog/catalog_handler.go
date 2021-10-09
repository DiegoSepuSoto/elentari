package catalog

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type catalogHandler struct {
	catalogUseCase usecase.CatalogUseCase
}

func NewCatalogHandler(e *echo.Echo, catalogUseCase usecase.CatalogUseCase) *catalogHandler {
	h := &catalogHandler{catalogUseCase: catalogUseCase}
	const basePath = "/v1/catalog"
	catalogGroup := e.Group(basePath)
	catalogGroup.GET("", h.getCatalog)

	return h
}

func (h *catalogHandler) getCatalog(c echo.Context) error {
	catalogPage, err := h.catalogUseCase.GetCatalogPage()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, catalogPage)
}