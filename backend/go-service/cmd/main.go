package main

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	routes.RegisterUserRoutes(r)
	routes.RegisterAccountRoutes(r)
	routes.RegisterTransactionRoutes(r)
	routes.RegisterCategoryRoutes(r)
	routes.RegisterTemplateRoutes(r)
	routes.RegisterSettingsRoutes(r)

	r.Run(":8080")
}
