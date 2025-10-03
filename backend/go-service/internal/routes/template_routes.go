package routes

import (
	"accountantapp/go-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTemplateRoutes(r *gin.Engine) {
	tmpls := r.Group("/templates")
	{
		tmpls.POST("/", controllers.CreateTemplate)
		tmpls.GET("/", controllers.GetAllTemplates)
		tmpls.GET("/:id", controllers.GetTemplateByID)
		tmpls.PUT("/:id", controllers.UpdateTemplate)
		tmpls.DELETE("/:id", controllers.DeleteTemplate)
	}
}
