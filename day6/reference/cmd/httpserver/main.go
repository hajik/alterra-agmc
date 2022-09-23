package main

import (
	"appday6/internal/core/service/gamesrv"
	"appday6/internal/handlers/gamehdl"
	"appday6/internal/repositories/gamesrepo"
	"appday6/pkg/uidgen"

	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}
