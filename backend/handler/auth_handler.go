package handler

import (
	"net/http"

	"github.com/fauzan264/user-appointments/auth"
	"github.com/fauzan264/user-appointments/helper"
	"github.com/fauzan264/user-appointments/middleware"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService auth.Service
	jwtService middleware.JWTService
}

func NewAuthHandler(authService auth.Service, jwtService middleware.JWTService) *authHandler {
	return &authHandler{authService, jwtService}
}

func (h *authHandler) RegisterUser(c *gin.Context) {
	var input auth.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		c.Set("error", errorMessage)
		response := helper.APIResponse(
			false,
			"Register account failed.",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.authService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		c.Set("error", errorMessage)
		response := helper.APIResponse(
			false,
			"Register account failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := auth.FormatAuth(newUser, "")
	response := helper.APIResponse(
		true,
		"Account has been registered",
		formatter,
	)
	c.JSON(http.StatusOK, response)
	return
}