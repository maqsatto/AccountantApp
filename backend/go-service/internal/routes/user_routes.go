package routes

import (
	"accountantapp/go-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.GetAllUsers)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
