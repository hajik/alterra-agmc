package routes

import (
	"test-github/controller"
	"test-github/util/logger"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// declare echo v4
	e := echo.New()
	//declare logger
	logger.DefaultLogger(e)
	// declare controller
	c := controller.Controller{}
	// create groups api
	apiGroup := e.Group("/api")
	c.ApiGroup(apiGroup, e)

	return e
}
