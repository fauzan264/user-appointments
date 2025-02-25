package user

import (
	"github.com/fauzan264/user-appointments/appointment"
	"github.com/google/uuid"
)

type User struct {
	ID 					uuid.UUID	`gorm:"type:char(36);primary_key"`	
	Name 				string		`gorm:"type:varchar(50);not null"`
	Username 			string		`gorm:"type:varchar(50);unique;not null"`
	PreferredTimeZone 	string		`gorm:"type:timestamp"`

	CreatorAppointment	[]appointment.Appointment
	AppointmentUsers	[]appointment.AppointmentUser
}