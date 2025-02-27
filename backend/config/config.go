package config

import (
	"log"
	"os"
	"strconv"

	"github.com/fauzan264/user-appointments/helper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey	string
	AppHost		string
	AppPort		string
	Debug		string
}

func LoadConfig() *Config {
	return &Config{
		SecretKey: os.Getenv("SECRET_KEY"),
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
		Debug: os.Getenv("DEBUG"),
	}
}

func InitLogger() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatal("Error creating logs folder: ", err)
		}
	}

	logFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	} else {
		log.SetOutput(logFile)
	}
}

func SetupGinMode(debug string) {
	mode, err := strconv.ParseBool(debug)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	if mode {
		gin.SetMode(gin.DebugMode)
	}
}

func init() {
	InitLogger()

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("gtfield", helper.ValidateStartBeforeEnd)
	}
}