package notification

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type notificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
}

func NewNotificationHandler(e *echo.Echo, notificationUseCase usecase.NotificationUseCase, jwtMiddleware echo.MiddlewareFunc) *notificationHandler {
	h := &notificationHandler{notificationUseCase: notificationUseCase}
	const basePath = "/v1/notification"
	notificationGroup := e.Group(basePath)
	notificationGroup.Use(jwtMiddleware)
	notificationGroup.GET("", h.getNotificationsPage)

	return h
}

// @Summary Pantalla de Notificaciones
// @Description Devuelve el contenido a mostrar en la pantalla de notificaciones de la aplicación
// @Tags BFF V1 - Pantalla de Notificaciones
// @Accept json
// @Produce json
// @Param Token-Autorización header string true "Bearer token"
// @Success 200 {object} models.NotificationsPage "OK"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 502 {object} map[string]interface{} "BadGateway"
// @Router /v1/notification [get]
func (h *notificationHandler) getNotificationsPage(c echo.Context) error {
	notificationsPage, err := h.notificationUseCase.GetNotifications()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, notificationsPage)
}
