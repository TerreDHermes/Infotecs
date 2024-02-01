package service

import "Infotecs/pkg/repository"

type CreateService struct {
	repo repository.Create
}

func NewCreateService(repo repository.Create) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateUser() (int, float64, error) {
	return s.repo.CreateUser()
}
