package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateWallet(c *gin.Context) {
	id, balance, err := h.services.Create.CreateUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"balance": balance,
	})

}
