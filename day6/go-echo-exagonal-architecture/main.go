package main

import (
	"fmt"
	"test-github/domain/repository"
	"test-github/routes"
)

func main() {
	// call database configuration->connection
	config := repository.Config{}
	c, _ := config.NewConfig()

	fmt.Println("Welcome to the webserver")
	e := routes.New()

	e.Logger.Fatal(e.Start(c.Port))
}