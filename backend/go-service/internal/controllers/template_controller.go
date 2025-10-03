package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTemplate(c *gin.Context) {
	var input models.Template
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tmpl, _ := services.CreateTemplateService(&input)
	c.JSON(http.StatusCreated, tmpl)
}

func GetAllTemplates(c *gin.Context) {
	tmpls, _ := services.GetAllTemplatesService()
	c.JSON(http.StatusOK, tmpls)
}

func GetTemplateByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tmpl, _ := services.GetTemplateByIDService(id)
	c.JSON(http.StatusOK, tmpl)
}

func UpdateTemplate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Template
	c.ShouldBindJSON(&input)
	tmpl, _ := services.UpdateTemplateService(id, &input)
	c.JSON(http.StatusOK, tmpl)
}

func DeleteTemplate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = services.DeleteTemplateService(id)
	c.JSON(http.StatusOK, gin.H{"message": "Template deleted"})
}
