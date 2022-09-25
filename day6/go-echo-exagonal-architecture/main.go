package main

import (
	"fmt"
	"test-github/domain/repository"
	"test-github/domain/repository/mongodb"
	"test-github/routes"
)

func main() {
	// call database configuration->connection
	config := repository.Config{}
	c, _ := config.NewConfig()

	// call mongodb
	mDB := mongodb.MongoDB{}
	mDB.NewDB()

	fmt.Println("Welcome to the webserver")
	e := routes.New()

	e.Logger.Fatal(e.Start(c.Port))
}
