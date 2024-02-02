package Infotecs

type Wallet struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type TransactionRequest struct {
	To     int     `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
