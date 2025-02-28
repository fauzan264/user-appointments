package appointment

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateAppointment(appointment Appointment) (Appointment, error)
	CreateAppointmentUser(appointmentUser AppointmentUser) (AppointmentUser, error)
	GetAppointmentByCreatorID(userID uuid.UUID) ([]Appointment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAppointment(appointment Appointment) (Appointment, error) {
	if err := r.db.Create(&appointment).Error; err != nil {
		return appointment, err
	}

	return appointment, nil
}

func (r *repository) CreateAppointmentUser(appointmentUser AppointmentUser) (AppointmentUser, error) {
	if err := r.db.Create(&appointmentUser).Error; err != nil {
		return appointmentUser, err
	}

	return appointmentUser, nil
}

func (r *repository) GetAppointmentByCreatorID(userID uuid.UUID) ([]Appointment, error) {
	appointments := []Appointment{}
	if err := r.db.Where("creator_id = ?", userID).Find(&appointments).Error; err != nil {
		return appointments, err
	}

	return appointments, nil
}