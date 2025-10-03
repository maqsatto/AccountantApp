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
	routes.RegisterTestRoutes(r)

	r.Run(":8080")
}
