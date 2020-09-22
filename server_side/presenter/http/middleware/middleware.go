package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitMiddleware ...
func InitMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
