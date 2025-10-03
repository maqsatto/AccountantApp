package routes

import (
	"accountantapp/go-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	cats := r.Group("/categories")
	{
		cats.POST("/", controllers.CreateCategory)
		cats.GET("/", controllers.GetAllCategories)
		cats.GET("/:id", controllers.GetCategoryByID)
		cats.PUT("/:id", controllers.UpdateCategory)
		cats.DELETE("/:id", controllers.DeleteCategory)
	}
}
