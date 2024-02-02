package handler

import (
	"Infotecs/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	wallets := router.Group("/api/v1/wallet")
	{
		wallets.POST("/", h.CreateWallet)

		walletsOperation := wallets.Group("/:walletId", h.WalletIdentity)
		{
			walletsOperation.POST("/send", h.Send)
			walletsOperation.GET("/history", h.TransactionsInfo)
			walletsOperation.GET("/", h.WalletInfo)
		}
	}

	router.POST("/api/v1/wallet", h.CreateWallet)
	return router
}
