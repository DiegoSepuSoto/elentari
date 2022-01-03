package post

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type postHandler struct {
	postUseCase usecase.PostUseCase
}

func NewPostHandler(e *echo.Echo, postUseCase usecase.PostUseCase, jwtMiddleware echo.MiddlewareFunc) *postHandler {
	h := &postHandler{postUseCase: postUseCase}
	const basePath = "/v1/post"
	postGroup := e.Group(basePath)
	postGroup.Use(jwtMiddleware)
	postGroup.GET("/:id", h.getPostPage)
	postGroup.GET("/search/:searchTerm", h.getPostsByTerm)

	return h
}

// @Summary Pantalla de Publicación
// @Description Devuelve el detalle de una publicación pora ser mostrado en la aplicación
// @Tags BFF V1 - Publicaciones
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Param id path string true "123456"
// @Success 200 {object} models.PostPage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/post/{id} [get]
func (h *postHandler) getPostPage(c echo.Context) error {
	postID := c.Param("id")
	postPage, err := h.postUseCase.GetPostPage(postID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, postPage)
}

// @Summary Resultados de Búsqueda
// @Description Devuelve todas las publicaciones que cumplen con el término de búsqueda
// @Tags BFF V1 - Publicaciones
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Param searchTerm path string true "bienestar"
// @Success 200 {object} models.PostSearchResult "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/post/search/{searchTerm} [get]
func (h *postHandler) getPostsByTerm(c echo.Context) error {
	searchTerm := c.Param("searchTerm")
	postsByTerm, err := h.postUseCase.GetPostsByTerm(searchTerm)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, postsByTerm)
}
