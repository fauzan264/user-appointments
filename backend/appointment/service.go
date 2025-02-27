package appointment

import (
	"errors"

	"github.com/fauzan264/user-appointments/helper"
	"github.com/fauzan264/user-appointments/user"
	"github.com/google/uuid"
)

type Service interface {
	CreateAppointment(input CreateAppointmentInput) (Appointment, error)
	CreateAppointmentUser(input CreateAppointmentUserInput) (AppointmentUser, error)
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

	// check Appointment time is working hours
	if !helper.IsWithinWorkingHours(input.Start, getUser.PreferredTimeZone) || !helper.IsWithinWorkingHours(input.End, getUser.PreferredTimeZone) {
		return Appointment{}, errors.New("Appointment time is outside working hours (08:00 - 17:00).")
	}

	appointmentUsers = append(appointmentUsers, appointmentUser)
	appointment.AppointmentUsers = appointmentUsers

	createAppointment, err := s.repository.CreateAppointment(appointment)
	if err != nil {
		return createAppointment, err
	}

	return createAppointment, nil
}

func(s *service) CreateAppointmentUser(input CreateAppointmentUserInput) (AppointmentUser, error) {
	appointmentUser := AppointmentUser{
		ID: uuid.New(),
		AppointmentID: input.AppointmentID,
		UserID: input.UserID,
	}

	getUser, err := s.userRepository.GetUserByid(input.UserID)
	if err != nil {
		return AppointmentUser{}, err
	}
	appointmentUser.User = getUser
	
	createAppointmentUser, err := s.repository.CreateAppointmentUser(appointmentUser)
	if err != nil {
		return createAppointmentUser, err
	}

	return createAppointmentUser, nil
}