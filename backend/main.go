package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/user-appointments/appointment"
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
	appointmentRepository := appointment.NewRepository(db)

	// Services
	authService := auth.NewService(userRepository)
	userService := user.NewService(userRepository)
	appointmentService := appointment.NewService(appointmentRepository, userRepository)
	jwtService := middleware.NewJWTService()

	// handlers
	authHandler := handler.NewAuthHandler(authService, jwtService)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService, jwtService)


	api := router.Group("/api/v1")
	// auth
	api.POST("/register", authHandler.RegisterUser)
	api.POST("/login", authHandler.Login)

	// appointment
	api.POST("/appointment", middleware.AuthMiddleware(jwtService, userService), appointmentHandler.CreateAppointment)
	

	err := router.Run(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort))
	if err != nil {
		log.Fatal("Failed to start the server", err.Error())
	}
}