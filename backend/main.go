package main

import (
	"fmt"
	"log"
	"time"

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Ganti dengan domain frontend
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour, // Cache preflight selama 12 jam
	}))

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
	api.POST("/auth/register", authHandler.RegisterUser)
	api.POST("/auth/login", authHandler.Login)
	api.GET("/auth/session", middleware.AuthMiddleware(jwtService, userService), authHandler.GetSession)

	// appointment
	api.POST("/appointment", middleware.AuthMiddleware(jwtService, userService), appointmentHandler.CreateAppointment)
	api.POST("/appointment_user", middleware.AuthMiddleware(jwtService, userService), appointmentHandler.CreateAppointmentUser)

	err := router.Run(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort))
	if err != nil {
		log.Fatal("Failed to start the server", err.Error())
	}
}