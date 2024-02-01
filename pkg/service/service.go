package service

import "Infotecs/pkg/repository"

type Wallet interface {
	CreateWallet() (int, float64, error)
	SearchId(walletId int) (int, float64, error)
}
type Service struct {
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: NewCreateService(repos.Wallet),
	}
}
