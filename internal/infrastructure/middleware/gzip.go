package middleware

import (
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func Gzip() echo.MiddlewareFunc {
	return emiddleware.Gzip()
}
