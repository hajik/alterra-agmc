package routes

import (
	"mvcapp/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserController)
	e.GET("/users/:id", controllers.DetailUserController)
	e.POST("/users", controllers.StoreUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}
