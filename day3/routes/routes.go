package routes

import (
	"mvcapp/controllers"
	"mvcapp/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// declare echo
	e := echo.New()

	// list routing
	// users
	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.StoreUserController)
	// book
	e.GET("/books", controllers.GetBookController)
	e.GET("/books/id", controllers.DetailBookController)

	r := e.Group("/JWT") // setting routing group

	// validation token using jwt
	r.Use(middleware.JWTWithConfig(middlewares.JWTConfig()))

	// book
	r.POST("/books", controllers.StoreBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	// users
	r.GET("/users", controllers.GetUserController)
	r.GET("/users/:id", controllers.DetailUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}
