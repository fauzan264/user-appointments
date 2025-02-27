package appointment

import (
	"time"

	"github.com/fauzan264/user-appointments/user"
	"github.com/google/uuid"
)

type CreateAppointmentInput struct {
	Title 			string 		`json:"title" binding:"required"`
	Start 			time.Time	`json:"start" binding:"required"`
	End				time.Time	`json:"end" binding:"required"`
	User			user.User
	AppointmentUser	AppointmentUser
}

type CreateAppointmentUserInput struct {
	AppointmentID	uuid.UUID 	`json:"appointment_id" binding:"required"`
	UserID			uuid.UUID	`json:"user_id" binding:"required"`
	User			user.User
}