package appointment

import "gorm.io/gorm"

type Repository interface {
	CreateAppointment(appointment Appointment) (Appointment, error)
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