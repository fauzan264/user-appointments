package user

import "github.com/google/uuid"

type Service interface {
	GetUserByID(id uuid.UUID) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserByID(id uuid.UUID) (User, error) {
	user, err := s.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}