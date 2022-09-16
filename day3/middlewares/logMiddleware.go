package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{${time_rfc3339_nano} ` +
			`${host}${uri} ${method} ${status} ${latency_human} ${error}` + "\n",
	}))
}
