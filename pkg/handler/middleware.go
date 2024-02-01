package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	walletIDCTX      = "walletId"
	walletBalanceCTX = "walletBalance"
)

func (h *Handler) WalletIdentity(c *gin.Context) {
	// Извлекаем ID кошелька из параметра запроса
	var walletBalance float64
	walletIDString := c.Param("walletId")
	if walletIDString == "" {
		newErrorResponse(c, http.StatusBadRequest, "id is empty")
		return
	}
	walletID, err := strconv.Atoi(walletIDString)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	walletID, walletBalance, err = h.services.Wallet.SearchId(walletID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.Set(walletIDCTX, walletID)
	c.Set(walletBalanceCTX, walletBalance)
	//c.Next()
}
