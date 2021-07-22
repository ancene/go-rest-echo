package middleware

import (
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return emiddleware.LoggerWithConfig(emiddleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}
