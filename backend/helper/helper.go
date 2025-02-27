package helper

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status		bool 		`json:"status"`
	Message 	string 		`json:"message"`
	Data		interface{}	`json:"data"`
}

func APIResponse(status bool, message string, data interface{}) Response {
	jsonResponse := Response{
		Status : status,
		Message: message,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ValidateStartBeforeEnd(fl validator.FieldLevel) bool {
	start := fl.Parent().FieldByName("Start").Interface().(time.Time)
	end := fl.Field().Interface().(time.Time)

	return !start.After(end)
}

func IsWithinWorkingHours(utcTime time.Time, tz string) bool {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Println("Error loading timezone:", err)
		return false
	}

	localTime := utcTime.In(loc)
	hour := localTime.Hour()

	return hour >= 8 && hour < 17
}