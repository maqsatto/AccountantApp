package routes

import (
	"accountantapp/go-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSettingsRoutes(r *gin.Engine) {
	settings := r.Group("/settings")
	{
		settings.POST("/", controllers.CreateSetting)
		settings.GET("/", controllers.GetAllSettings)
		settings.GET("/:id", controllers.GetSettingByID)
		settings.PUT("/:id", controllers.UpdateSetting)
		settings.DELETE("/:id", controllers.DeleteSetting)
	}
}
