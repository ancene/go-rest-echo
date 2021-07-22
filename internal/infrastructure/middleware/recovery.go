package middleware

import (
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func Recovery() echo.MiddlewareFunc {
	return emiddleware.Recover()
}
