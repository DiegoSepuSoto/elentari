package middleware

import (
	"elentari/src/application/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type jwtValidationMiddleware struct {
	authUsecase usecase.AuthUseCase
}

func NewJWTMiddleware(authUseCase usecase.AuthUseCase) echo.MiddlewareFunc {
	jwtMiddleware := &jwtValidationMiddleware{authUsecase: authUseCase}
	return jwtMiddleware.jwtValidationMiddlewareFunc()
}

func (m *jwtValidationMiddleware) jwtValidationMiddlewareFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(echo.HeaderAuthorization)
			isValid, err := m.authUsecase.ValidateToken(token)
			if err != nil || !isValid {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
			}
			err = next(c)
			return err
		}
	}
}
