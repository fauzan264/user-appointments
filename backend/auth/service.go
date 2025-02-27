package auth

import "os/user"

type Service interface {
	RegisterUser(input RegisterUserInput) (user.User, error)
}