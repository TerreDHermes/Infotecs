package service

import (
	"Infotecs"
	"Infotecs/pkg/repository"
)

type Wallet interface {
	CreateWallet() (int, float64, error)
	SearchId(walletId int) (int, float64, error)
	Send(walletFromId int, walletToId int, Amount float64) error
	TransactionsInfo(walletId int) ([]Infotecs.TransactionInfo, error)
}
type Service struct {
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: NewCreateService(repos.Wallet),
	}
}
