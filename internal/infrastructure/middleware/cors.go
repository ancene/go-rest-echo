package middleware

import (
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	return emiddleware.CORS()
}
