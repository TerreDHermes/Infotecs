package repository

import (
	"Infotecs"
	"github.com/jmoiron/sqlx"
)

type Wallet interface {
	CreateWallet() (int, float64, error)
	SearchId(walletId int) (int, float64, error)
	Send(walletFromId int, walletToId int, Amount float64) error
	TransactionsInfo(walletId int) ([]Infotecs.TransactionInfo, error)
}

type Repository struct {
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet: NewCreatePostgres(db),
	}
}
