package appointment

import (
	"time"

	"github.com/fauzan264/user-appointments/user"
)

type CreateAppointmentInput struct {
	Title 			string 		`json:"title" binding:"required"`
	Start 			time.Time	`json:"start" binding:"required"`
	End				time.Time	`json:"end" binding:"required"`
	User			user.User
	AppointmentUser	AppointmentUser
}