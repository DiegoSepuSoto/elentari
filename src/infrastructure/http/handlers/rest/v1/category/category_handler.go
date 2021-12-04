package category

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type categoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(e *echo.Echo, categoryUseCase usecase.CategoryUseCase, jwtMiddleware echo.MiddlewareFunc) *categoryHandler {
	h := &categoryHandler{categoryUseCase: categoryUseCase}
	const basePath = "/v1/category"
	categoryGroup := e.Group(basePath)
	categoryGroup.Use(jwtMiddleware)
	categoryGroup.GET("/:id/posts", h.getCategoryPostsPage)

	return h
}

// @Summary Publicaciones de una categoría
// @Description Devuelve todas las publicaciones que presentan la categoría enviada por parámetro
// @Tags BFF V1 - Categorías
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Param id path string true "123456"
// @Success 200 {object} models.CategoryPostsPage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/category/{id}/posts [get]
func (h *categoryHandler) getCategoryPostsPage(c echo.Context) error {
	categoryID := c.Param("id")
	categoryPosts, err := h.categoryUseCase.GetCategoryPostsPage(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, categoryPosts)
}