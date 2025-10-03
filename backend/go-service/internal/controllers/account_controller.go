package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var input models.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, _ := services.CreateAccountService(&input)
	c.JSON(http.StatusCreated, account)
}

func GetAllAccounts(c *gin.Context) {
	accounts, _ := services.GetAllAccountsService()
	c.JSON(http.StatusOK, accounts)
}

func GetAccountByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	account, _ := services.GetAccountByIDService(id)
	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Account
	c.ShouldBindJSON(&input)
	account, _ := services.UpdateAccountService(id, &input)
	c.JSON(http.StatusOK, account)
}

func DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = services.DeleteAccountService(id)
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}
