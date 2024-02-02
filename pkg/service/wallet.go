package service

import (
	"Infotecs"
	"Infotecs/pkg/repository"
)

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

func (s *CreateService) Send(walletFromId int, walletToId int, Amount float64) error {
	return s.repo.Send(walletFromId, walletToId, Amount)
}

func (s *CreateService) TransactionsInfo(walletId int) ([]Infotecs.TransactionInfo, error) {
	return s.repo.TransactionsInfo(walletId)
}
