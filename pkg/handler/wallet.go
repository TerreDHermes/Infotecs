package handler

import (
	"Infotecs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateWallet(c *gin.Context) {
	id, balance, err := h.services.Wallet.CreateWallet()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"balance": balance,
	})

}

func (h *Handler) WalletInfo(c *gin.Context) {
	walletId, _ := c.Get(walletIDCTX)
	walletBalance, _ := c.Get(walletBalanceCTX)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      walletId,
		"balance": walletBalance,
	})
}

func (h *Handler) Send(c *gin.Context) {
	var transaction Infotecs.TransactionRequest
	walletFromId, _ := c.Get(walletIDCTX)
	walletFromBalance, _ := c.Get(walletBalanceCTX)
	if err := c.ShouldBindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// проверяем, что получатель существует
	walletToId, _, err := h.services.Wallet.SearchId(transaction.To)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// проверяем, что на балансе достаточно средств
	if walletFromBalance.(float64)-transaction.Amount < 0 {
		newErrorResponse(c, http.StatusBadRequest, "The amount is not enough")
		return
	}
	//проверяем, чтобы сумма перевода не была отрицательным числом
	if transaction.Amount < 0 {
		newErrorResponse(c, http.StatusBadRequest, "The transfer amount cannot be a negative number")
		return
	}

	// отправка средств
	if err := h.services.Wallet.Send(walletFromId.(int), walletToId, transaction.Amount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "The transaction was successful",
	})
	//walletId, _ := c.Get(walletIDCTX)
	//walletBalance, _ := c.Get(walletBalanceCTX)
}
