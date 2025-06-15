package routes

import (
	"github.com/chinotanismail/internal-transfer-system/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/accounts", handlers.CreateAccount)
	r.GET("/accounts/:account_id", handlers.GetAccount)
	r.POST("/transactions", handlers.CreateTransaction)

	return r
}
