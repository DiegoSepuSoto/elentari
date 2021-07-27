package post

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type postHandler struct {
	postUseCase usecase.PostUseCase
}

func NewPostHandler(e *echo.Echo, postUseCase usecase.PostUseCase) *postHandler {
	h := &postHandler{postUseCase: postUseCase}
	const basePath = "/v1/posts"
	postGroup := e.Group(basePath)
	postGroup.GET("", h.getPosts)
	postGroup.GET("/:id", h.getPostsByServiceID)

	return h
}

func (h postHandler) getPosts(c echo.Context) error {
	posts, err := h.postUseCase.GetPosts()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, posts)
}

func (h postHandler) getPostsByServiceID(c echo.Context) error {
	serviceID := c.Param("id")

	posts, err := h.postUseCase.GetPostsByServiceID(serviceID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}

	return c.JSON(http.StatusOK, posts)
}