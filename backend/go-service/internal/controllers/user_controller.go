package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Accounts = []models.Account{}
	input.Transactions = []models.Transaction{}
	input.Categories = []models.Category{}
	input.Templates = []models.Template{}
	input.Settings = []models.AppSettings{}

	user, err := services.CreateUserService(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	users, _ := services.GetAllUsersService()
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.UpdateUserService(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = services.DeleteUserService(id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
