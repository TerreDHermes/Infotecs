package repository

import (
	"Infotecs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
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
	timeFormatted, err2 := getCurrentTimeFormatted("Europe/Moscow")
	if err2 != nil {
		return err2
	}
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
	query := "INSERT INTO transactions (time, wallet_from_id, wallet_to_id, amount) VALUES ($1, $2, $3, $4)"
	tx.QueryRow(query, timeFormatted, walletFromId, walletToId, Amount)

	// Фиксация транзакции
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *WalletPostgres) TransactionsInfo(walletId int) ([]Infotecs.TransactionInfo, error) {
	// Запрос истории транзакций
	rows, err := r.db.Query(`
        SELECT time, wallet_from_id, wallet_to_id, amount
        FROM transactions
        WHERE wallet_from_id = $1 OR wallet_to_id = $1
        ORDER BY time DESC
    `, walletId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []Infotecs.TransactionInfo
	var timeUTC time.Time
	for rows.Next() {
		var transaction Infotecs.TransactionInfo
		if err := rows.Scan(&timeUTC, &transaction.From, &transaction.To, &transaction.Amount); err != nil {
			return nil, err
		}
		// Приводим время к нужному часовому поясу
		location, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			return nil, err
		}
		transaction.Time = timeUTC.In(location).Format(time.RFC3339)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func getCurrentTimeFormatted(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	return time.Now().In(location).Format(time.RFC3339), nil
}
