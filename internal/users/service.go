package users

import (
	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
	Store(edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.User, error) {
	u, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Store(edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, err
	}

	lastId++

	user, err := s.repository.Store(lastId, edad, nombre, apellido, email, fecha_creacion, altura, activo)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
