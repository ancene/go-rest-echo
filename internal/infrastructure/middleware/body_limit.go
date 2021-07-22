package middleware

import (
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func BodyLimit() echo.MiddlewareFunc {
	return emiddleware.BodyLimit("2M")
}
