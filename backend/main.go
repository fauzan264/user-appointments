package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/user-appointments/auth"
	"github.com/fauzan264/user-appointments/config"
	"github.com/fauzan264/user-appointments/handler"
	"github.com/fauzan264/user-appointments/middleware"
	"github.com/fauzan264/user-appointments/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDatabase(&gin.Context{})
	cfg := config.LoadConfig()

	config.SetupGinMode(cfg.Debug)

	router := gin.New()
	router.Use(cors.Default())

	// Repositories
	userRepository := user.NewRepository(db)

	// Services
	authService := auth.NewService(userRepository)
	jwtService := middleware.NewJWTService()

	// handlers
	authHandler := handler.NewAuthHandler(authService, jwtService)

	api := router.Group("/api/v1")
	api.POST("/register", authHandler.RegisterUser)

	err := router.Run(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort))
	if err != nil {
		log.Fatal("Failed to start the server", err.Error())
	}
}