package repository

import (
	"Infotecs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *WalletPostgres) Send(walletFromId int, walletToId int, Amount float64) error {
	// Выполнение транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Print("Failed to rollback transaction:", err)
		}
	}()

	// Уменьшение баланса у отправителя
	_, err = tx.Exec("UPDATE wallets SET balance = balance - $1 WHERE id = $2", Amount, walletFromId)
	if err != nil {
		return err
	}

	// Увеличение баланса у получателя
	_, err = tx.Exec("UPDATE wallets SET balance = balance + $1 WHERE id = $2", Amount, walletToId)
	if err != nil {
		return err
	}

	// Запись транзакции в таблицу transactions
	query := "INSERT INTO transactions (wallet_from_id, wallet_to_id, amount) VALUES ($1, $2, $3)"
	tx.QueryRow(query, walletFromId, walletToId, Amount)

	// Фиксация транзакции
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
