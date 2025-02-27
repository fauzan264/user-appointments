package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID 					uuid.UUID	`gorm:"column:id;type:char(36);primary_key"`	
	Name 				string		`gorm:"column:name;type:varchar(50);not null"`
	Username 			string		`gorm:"column:username;type:varchar(50);uniqueIndex:idx_username;not null"`
	Password			string		`gorm:"column:password;type:varchar(200);not null"`
	PreferredTimeZone 	string		`gorm:"column:preferred_timezone;type:varchar(100)"`
}