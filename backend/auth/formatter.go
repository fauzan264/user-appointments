package auth

import (
	"github.com/fauzan264/user-appointments/user"
	"github.com/google/uuid"
)


type AuthFormatter struct {
	ID					uuid.UUID	`json:"id"`
	Name				string		`json:"name"`
	Username			string		`json:"username"`
	PreferredTimeZone	string		`json:"preferred_timezone"`
	Token				string		`json:"token,omitempty"`
}

func FormatAuth(auth user.User, token string) AuthFormatter {
	formatter := AuthFormatter{
		ID: auth.ID,
		Name: auth.Name,
		Username: auth.Username,
		PreferredTimeZone: auth.PreferredTimeZone,
	}

	return formatter
}

