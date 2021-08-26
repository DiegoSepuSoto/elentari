package category

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type categoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(e *echo.Echo, categoryUseCase usecase.CategoryUseCase) *categoryHandler {
	h := &categoryHandler{categoryUseCase: categoryUseCase}
	const basePath = "/v1/category"
	categoryGroup := e.Group(basePath)
	categoryGroup.GET("/:id/posts", h.getCategoryPostsPage)

	return h
}

func (h *categoryHandler) getCategoryPostsPage(c echo.Context) error {
	categoryID := c.Param("id")
	categoryPosts, err := h.categoryUseCase.GetCategoryPostsPage(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, categoryPosts)
}