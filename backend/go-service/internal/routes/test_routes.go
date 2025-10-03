package routes

import (
	"accountantapp/go-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTestRoutes(r *gin.Engine) {
	r.POST("/test-user", controllers.TestCreateUser)
}
