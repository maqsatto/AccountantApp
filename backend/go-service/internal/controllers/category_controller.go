package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat, _ := services.CreateCategoryService(&input)
	c.JSON(http.StatusCreated, cat)
}

func GetAllCategories(c *gin.Context) {
	cats, _ := services.GetAllCategoriesService()
	c.JSON(http.StatusOK, cats)
}

func GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, _ := services.GetCategoryByIDService(id)
	c.JSON(http.StatusOK, cat)
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Category
	c.ShouldBindJSON(&input)
	cat, _ := services.UpdateCategoryService(id, &input)
	c.JSON(http.StatusOK, cat)
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = services.DeleteCategoryService(id)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
