package main

import(
	"mvcapp/config"
	"mvcapp/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
