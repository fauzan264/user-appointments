package appointment

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentFormatter struct {
	ID 							uuid.UUID 	`json:"id"`
	Title						string		`json:"title"`
	CreatorID					uuid.UUID	`json:"creator_id"`
	Start						time.Time	`json:"start"`
	End							time.Time	`json:"end"`
	AppointmentUsersFormatter	[]AppointmentUserFormatter
}

type AppointmentUserFormatter struct {
	ID				uuid.UUID	`json:"id"`
	Name			string		`json:"name"`
	Username 		string		`json:"username"`
}

func FormatAppointmentUser(appointmentUser AppointmentUser) AppointmentUserFormatter {
	formatter := AppointmentUserFormatter{
		ID: appointmentUser.ID,
		Name: appointmentUser.User.Name,
		Username: appointmentUser.User.Username,
	}

	return formatter
}

func FormatAppointmentUsers(appointmentUsers []AppointmentUser) []AppointmentUserFormatter {
	appointmentUsersFormatter := []AppointmentUserFormatter{}

	for _, appointmentUser := range appointmentUsers {
		appointmentUser := FormatAppointmentUser(appointmentUser)
		appointmentUsersFormatter = append(appointmentUsersFormatter, appointmentUser)
	}

	return appointmentUsersFormatter
}

func FormatAppointment(appointment Appointment) AppointmentFormatter {
	formatter := AppointmentFormatter{
		ID: appointment.ID,
		Title: appointment.Title,
		CreatorID: appointment.CreatorID,
		Start: appointment.Start,
		End: appointment.End,
	}

	appointmentUsers := FormatAppointmentUsers(appointment.AppointmentUsers)
	formatter.AppointmentUsersFormatter = appointmentUsers

	return formatter
}