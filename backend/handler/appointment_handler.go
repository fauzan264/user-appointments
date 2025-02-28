package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/fauzan264/user-appointments/appointment"
	"github.com/fauzan264/user-appointments/helper"
	"github.com/fauzan264/user-appointments/middleware"
	"github.com/fauzan264/user-appointments/user"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	appointmentService appointment.Service
	jwtService middleware.JWTService
}

func NewAppointmentHandler(appointmentService appointment.Service, jwtService middleware.JWTService) *appointmentHandler {
	return &appointmentHandler{appointmentService, jwtService}
}

func (h *appointmentHandler) CreateAppointment(c *gin.Context) {
	var input appointment.CreateAppointmentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		if timeErr, ok := err.(*time.ParseError); ok {
			errorMessage := gin.H{"error": timeErr}
			
			response := helper.APIResponse(
				false,
				"Invalid time format",
				errorMessage,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			false,
			"Failed create appointment",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newAppointment, err := h.appointmentService.CreateAppointment(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(
			false,
			"Failed create appointment",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	response := helper.APIResponse(
		true,
		"Success create appointment",
		appointment.FormatAppointment(newAppointment),
	)
	c.JSON(http.StatusOK, response)
	return
}

func (h *appointmentHandler) CreateAppointmentUser(c *gin.Context) {
	var input appointment.CreateAppointmentUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		if timeErr, ok := err.(*time.ParseError); ok {
			errorMessage := gin.H{"error": timeErr}
			
			response := helper.APIResponse(
				false,
				"Invalid time format",
				errorMessage,
			)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			false,
			"Failed create appointment user",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newAppointment, err := h.appointmentService.CreateAppointmentUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(
			false,
			"Failed create appointment user",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	response := helper.APIResponse(
		true,
		"Success create appointment user",
		appointment.FormatAppointmentUser(newAppointment),
	)
	c.JSON(http.StatusOK, response)
	return
}

func (h *appointmentHandler) GetUserAppointments(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	log.Println(userID)

	getApppointments, err := h.appointmentService.GetAppointmentByCreatorID(userID)
	if err != nil {
		response := helper.APIResponse(false, "Failed to get appointments", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(true, "Success to get appointments", appointment.FormatAppointments(getApppointments))
	c.JSON(http.StatusOK, response)
	return
}