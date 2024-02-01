package service

import "Infotecs/pkg/repository"

type CreateService struct {
	repo repository.Wallet
}

func NewCreateService(repo repository.Wallet) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateWallet() (int, float64, error) {
	return s.repo.CreateWallet()
}

func (s *CreateService) SearchId(walletId int) (int, float64, error) {
	return s.repo.SearchId(walletId)
}
