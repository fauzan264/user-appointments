package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fauzan264/user-appointments/appointment"
	"github.com/fauzan264/user-appointments/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(c *gin.Context) *gorm.DB {
	var(
		DBHost		= os.Getenv("DB_HOST")
		DBPort		= os.Getenv("DB_PORT")
		DBName		= os.Getenv("DB_NAME")
		DBUser		= os.Getenv("DB_USER")
		DBPassword	= os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.Set("error", errorMessage)
		log.Fatal("Error: ", err.Error())
	}
	
	err = db.AutoMigrate(&user.User{}, &appointment.Appointment{}, &appointment.AppointmentUser{})
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.Set("error", errorMessage)
		log.Fatal("Error: ", err.Error())
	}

	return db
}