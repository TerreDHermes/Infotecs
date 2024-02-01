package repository

import (
	"Infotecs"
	"github.com/jmoiron/sqlx"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewCreatePostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) CreateWallet() (int, float64, error) {
	query := "INSERT INTO wallets (balance) VALUES ($1) RETURNING id, balance"
	row := r.db.QueryRow(query, 100.0)

	var wallet Infotecs.Wallet
	if err := row.Scan(&wallet.ID, &wallet.Balance); err != nil {
		return 0, 0, err
	}
	return wallet.ID, wallet.Balance, nil
}

func (r *WalletPostgres) SearchId(walletId int) (int, float64, error) {
	var wallet Infotecs.Wallet
	query := "SELECT id, balance FROM wallets WHERE id = $1"
	err := r.db.QueryRow(query, walletId).Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		return 0, 0, err
	}
	return wallet.ID, wallet.Balance, nil
}
