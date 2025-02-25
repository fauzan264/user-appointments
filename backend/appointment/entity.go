package appointment

import (
	"time"

	"github.com/fauzan264/user-appointments/user"
	"github.com/google/uuid"
)

type Appointment struct {
	ID 						uuid.UUID			`gorm:"column:id;type:char(36);primary_key"`
	Title 					string				`gorm:"column:title;type:varchar(50)"`
	CreatorID 				uuid.UUID			`gorm:"column:creator_id;type:char(36)"`
	Start 					time.Time			`gorm:"colunn:start;type:timestamp"`
	End 					time.Time			`gorm:"column:end;type:timestamp"`

	AppointmentUsers		[]AppointmentUser	`gorm:"foreignKey:AppointmentID"`
	User					user.User			`gorm:"foreignKey:CreatorID;references:ID"`
}

type AppointmentUser struct {
	ID				uuid.UUID			`gorm:"column:id;type:char(36);primary_key"`
	AppointmentID	uuid.UUID			`gorm:"column:appointment_id;type:char(36);uniqueIndex:idx_appointment_user"`
	UserID			uuid.UUID			`gorm:"column:user_id;type:char(36);uniqueIndex:idx_appointment_user"`

	Appointment		Appointment			`gorm:"foreignKey:AppointmentID;references:ID"`
	User			user.User			`gorm:"foreignKey:UserID;references:ID"`
}