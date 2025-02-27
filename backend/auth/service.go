package auth

import (
	"github.com/fauzan264/user-appointments/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput) (user.User, error)
}

type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (user.User, error) {
	user := user.User{
		ID: uuid.New(),
		Name: input.Name,
		Username: input.Username,
		PreferredTimeZone: input.PreferredTimeZone,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	createUser, err := s.repository.RegisterUser(user)
	if err != nil {
		return user, err
	}

	return createUser, nil
}