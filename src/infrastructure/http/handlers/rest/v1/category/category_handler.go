package category

import (
	"elentari/src/application/usecase"
	"elentari/src/application/usecase/category"
	"github.com/labstack/echo"
	"net/http"
)

type categoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(e *echo.Echo, useCase *category.UseCase) *categoryHandler {
	h := &categoryHandler{categoryUseCase: useCase}
	const basePath = "/v1/categories"
	postGroup := e.Group(basePath)
	postGroup.GET("", h.getCategories)

	return h
}

func (h categoryHandler) getCategories(c echo.Context) error {
	categories, err := h.categoryUseCase.GetCategories()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, categories)
}
