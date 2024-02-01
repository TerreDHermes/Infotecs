package handler

import (
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
