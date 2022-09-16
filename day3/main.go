package main

import (
	"mvcapp/config"
	"mvcapp/middlewares"
	"mvcapp/routes"
)

func main() {
	// call database
	config.InitDB()
	// call routes
	e := routes.New()
	// implement logger after route
	middlewares.LogMiddlewares(e)
	// declare echo start using port
	e.Logger.Fatal(e.Start(":8080"))
}
