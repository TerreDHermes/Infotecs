package repository

import (
	"Infotecs"
	"github.com/jmoiron/sqlx"
)

type CreateWalletPostgres struct {
	db *sqlx.DB
}

func NewCreatePostgres(db *sqlx.DB) *CreateWalletPostgres {
	return &CreateWalletPostgres{db: db}
}

func (r *CreateWalletPostgres) CreateUser() (int, float64, error) {
	query := "INSERT INTO wallets (balance) VALUES ($1) RETURNING id, balance"
	row := r.db.QueryRow(query, 100.0)

	var wallet Infotecs.Wallet
	if err := row.Scan(&wallet.ID, &wallet.Balance); err != nil {
		return 0, 0, err
	}
	return wallet.ID, wallet.Balance, nil
}
