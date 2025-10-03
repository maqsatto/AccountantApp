package routes

import (
	"accountantapp/go-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(r *gin.Engine) {
	accounts := r.Group("/accounts")
	{
		accounts.POST("/", controllers.CreateAccount)
		accounts.GET("/", controllers.GetAllAccounts)
		accounts.GET("/:id", controllers.GetAccountByID)
		accounts.PUT("/:id", controllers.UpdateAccount)
		accounts.DELETE("/:id", controllers.DeleteAccount)
	}
}
