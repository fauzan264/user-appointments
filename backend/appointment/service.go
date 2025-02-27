package appointment

import (
	"github.com/fauzan264/user-appointments/user"
	"github.com/google/uuid"
)

type Service interface {
	CreateAppointment(input CreateAppointmentInput) (Appointment, error)
}

type service struct {
	repository 		Repository
	userRepository	user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *service {
	return &service{repository, userRepository}
}

func(s *service) CreateAppointment(input CreateAppointmentInput) (Appointment, error) {
	appointmentID := uuid.New()
	
	appointment := Appointment{
		ID: appointmentID,
		Title: input.Title,
		CreatorID: input.User.ID,
		Start: input.Start,
		End: input.End,
	}

	appointmentUsers := []AppointmentUser{}
	appointmentUser := AppointmentUser{
		ID: uuid.New(),
		AppointmentID: appointmentID,
		UserID: input.User.ID,
	}

	getUser, err := s.userRepository.GetUserByid(input.User.ID)
	if err != nil {
		return Appointment{}, err
	}
	appointmentUser.User = getUser

	appointmentUsers = append(appointmentUsers, appointmentUser)
	appointment.AppointmentUsers = appointmentUsers

	createAppointment, err := s.repository.CreateAppointment(appointment)
	if err != nil {
		return createAppointment, err
	}

	return createAppointment, nil
}