package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
)

func CreateTemplateService(tmpl *models.Template) (*models.Template, error) {
	database.DB.Create(tmpl)
	return tmpl, nil
}

func GetAllTemplatesService() ([]models.Template, error) {
	var tmpls []models.Template
	database.DB.Find(&tmpls)
	return tmpls, nil
}

func GetTemplateByIDService(id int) (*models.Template, error) {
	var tmpl models.Template
	database.DB.First(&tmpl, id)
	return &tmpl, nil
}

func UpdateTemplateService(id int, input *models.Template) (*models.Template, error) {
	var tmpl models.Template
	database.DB.First(&tmpl, id)
	tmpl.Name = input.Name
	tmpl.Amount = input.Amount
	tmpl.CategoryID = input.CategoryID
	database.DB.Save(&tmpl)
	return &tmpl, nil
}

func DeleteTemplateService(id int) error {
	return database.DB.Delete(&models.Template{}, id).Error
}
