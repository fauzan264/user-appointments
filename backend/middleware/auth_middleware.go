package middleware

import (
	"net/http"
	"strings"

	"github.com/fauzan264/user-appointments/helper"
	"github.com/fauzan264/user-appointments/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func sendUnauthorizedResponse(c *gin.Context) {
	response := helper.APIResponse(
		false,
		"Unauthorized",
		nil,
	)
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func AuthMiddleware(jwtService JWTService, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		
		if !strings.Contains(authHeader, "Bearer") {
			sendUnauthorizedResponse(c)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			sendUnauthorizedResponse(c)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			sendUnauthorizedResponse(c)
			return
		}

		userID, err := uuid.Parse(claim["sub"].(string))
		if err != nil {
			sendUnauthorizedResponse(c)
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			sendUnauthorizedResponse(c)
			return
		}

		c.Set("currentUser", user)
	}
}