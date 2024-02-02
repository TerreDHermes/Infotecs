package Infotecs

type Wallet struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type TransactionRequest struct {
	To     int     `json:"to" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type TransactionInfo struct {
	Time   string  `json:"time"`
	From   int     `json:"from"`
	To     int     `json:"to"`
	Amount float64 `json:"amount"`
}
