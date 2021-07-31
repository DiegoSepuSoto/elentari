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
	const basePath = "/v1/post"
	postGroup := e.Group(basePath)
	postGroup.GET("/:id", h.getPostPage)

	return h
}

func (h postHandler) getPostPage(c echo.Context) error {
	postID := c.Param("id")
	postPage, err := h.postUseCase.GetPostPage(postID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, postPage)
}
