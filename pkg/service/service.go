package service

import "Infotecs/pkg/repository"

type Create interface {
	CreateUser() (int, float64, error)
}
type Service struct {
	Create
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Create: NewCreateService(repos.Create),
	}
}
