package controllers

import (
	"accountantapp/go-service/internal/models"
	"accountantapp/go-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSetting(c *gin.Context) {
	var input models.AppSettings
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	setting, _ := services.CreateSettingService(&input)
	c.JSON(http.StatusCreated, setting)
}

func GetAllSettings(c *gin.Context) {
	settings, _ := services.GetAllSettingsService()
	c.JSON(http.StatusOK, settings)
}

func GetSettingByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	setting, _ := services.GetSettingByIDService(id)
	c.JSON(http.StatusOK, setting)
}

func UpdateSetting(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.AppSettings
	c.ShouldBindJSON(&input)
	setting, _ := services.UpdateSettingService(id, &input)
	c.JSON(http.StatusOK, setting)
}

func DeleteSetting(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = services.DeleteSettingService(id)
	c.JSON(http.StatusOK, gin.H{"message": "Setting deleted"})
}
