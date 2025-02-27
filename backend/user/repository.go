package user

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			errMessage := errors.New("Username is already taken, please choose another one.")
			return user, errMessage
		}

		errMessage := errors.New("Internal Server Error")
		return user, errMessage
	}

	return user, nil
}

func (r *repository) GetUserByid(id uuid.UUID) (User, error) {
	var user User
	result := r.db.Where("id = ?", id).Find(&user)

	if result.RowsAffected == 0 {
		return user, errors.New("User not found")
	}

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}