package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := services.CreateTransaction(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tx)
}

func GetAllTransactions(c *gin.Context) {
	txs, err := services.GetAllTransactionsService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, txs)
}

func GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tx, err := services.GetTransactionByIDService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func UpdateTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := services.UpdateTransaction(id, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteTransaction(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}
