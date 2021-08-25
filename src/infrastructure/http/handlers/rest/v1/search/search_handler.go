package search

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type searchHandler struct {
	searchUseCase usecase.SearchUseCase
}

func NewSearchHandler(e *echo.Echo, searchUseCase usecase.SearchUseCase) *searchHandler {
	h := &searchHandler{searchUseCase: searchUseCase}
	const basePath = "/v1/search"
	searchGroup := e.Group(basePath)
	searchGroup.GET("", h.getSearch)

	return h
}

func (h *searchHandler) getSearch(c echo.Context) error {
	searchPage, err := h.searchUseCase.GetSearchPage()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, searchPage)
}