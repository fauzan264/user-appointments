package appointment

import (
	"os/user"
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID 						uuid.UUID			`gorm:"type:char(36);primary_key"`
	Title 					string				`gorm:"type:varchar(50)"`
	CreatorID 				uuid.UUID			`gorm:"type:char(36)"`
	Start 					time.Time			`gorm:"type:timestamp"`
	End 					time.Time			`gorm:"type:timestamp"`

	AppointmentUsers		[]AppointmentUser
	User					user.User			`gorm:"foreignKey:CreatorID;references:ID"`
}

type AppointmentUser struct {
	ID				uuid.UUID			`gorm:"type:char(36);primary_key"`
	AppointmentID	uuid.UUID			`gorm:"type:char(36);uniqueIndex:idx_appointment_user"`
	UserID			uuid.UUID			`gorm:"type:char(36);uniqueIndex:idx_appointment_user"`

	Appointment		Appointment			`gorm:"foreignKey:AppointmentID;references:ID"`
	User			user.User			`gorm:"foreignKey:UserID;references:ID"`
}