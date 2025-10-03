package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
)

func CreateSettingService(s *models.AppSettings) (*models.AppSettings, error) {
	database.DB.Create(s)
	return s, nil
}

func GetAllSettingsService() ([]models.AppSettings, error) {
	var s []models.AppSettings
	database.DB.Find(&s)
	return s, nil
}

func GetSettingByIDService(id int) (*models.AppSettings, error) {
	var s models.AppSettings
	database.DB.First(&s, id)
	return &s, nil
}

func UpdateSettingService(id int, input *models.AppSettings) (*models.AppSettings, error) {
	var s models.AppSettings
	database.DB.First(&s, id)
	s.Key = input.Key
	s.Value = input.Value
	database.DB.Save(&s)
	return &s, nil
}

func DeleteSettingService(id int) error {
	return database.DB.Delete(&models.AppSettings{}, id).Error
}
