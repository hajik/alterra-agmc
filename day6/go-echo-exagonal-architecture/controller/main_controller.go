package controller

import (
	"test-github/domain/service"
	"test-github/domain/service/society"

	m "test-github/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	society *society.SocietyService
	user    *service.UserService
}

func (c *Controller) NewsocietyController() {
	c.society = &society.SocietyService{}
	c.society.NewsocietyOrm()
}

func (c *Controller) NewUserController() {
	c.user = &service.UserService{}
	c.user.NewUserOrm()
}

func (c *Controller) ApiGroup(g *echo.Group, e *echo.Echo) {
	c.NewUserController()

	// auth
	g.POST("/login", c.user.Login)
	//society
	c.NewsocietyController()
	g.GET("/societies", c.society.ListAll, middleware.JWTWithConfig(m.JWTConfig()))

	// user
	g.Use(middleware.JWTWithConfig(m.JWTConfig()))
	g.GET("/user/list", c.user.ListAll)
	g.POST("/user/store/one", c.user.StoreOne)
	g.PUT("/user/update/:name", c.user.Update)
	g.DELETE("/user/delete/:name", c.user.Delete)
}
