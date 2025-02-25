package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/user-appointments/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = config.InitDatabase(&gin.Context{})
	cfg := config.LoadConfig()

	config.SetupGinMode(cfg.Debug)

	router := gin.New()
	router.Use(cors.Default())

	err := router.Run(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort))
	if err != nil {
		log.Fatal("Failed to start the server", err.Error())
	}
}