package routes

import (
	"accountantapp/go-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(r *gin.Engine) {
	txs := r.Group("/transactions")
	{
		txs.POST("/", controllers.CreateTransaction)
		txs.GET("/", controllers.GetAllTransactions)
		txs.GET("/:id", controllers.GetTransactionByID)
		txs.PUT("/:id", controllers.UpdateTransaction)
		txs.DELETE("/:id", controllers.DeleteTransaction)
	}
}
