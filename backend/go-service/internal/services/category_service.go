package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
)

func CreateCategoryService(cat *models.Category) (*models.Category, error) {
	database.DB.Create(cat)
	return cat, nil
}

func GetAllCategoriesService() ([]models.Category, error) {
	var cats []models.Category
	database.DB.Find(&cats)
	return cats, nil
}

func GetCategoryByIDService(id int) (*models.Category, error) {
	var cat models.Category
	database.DB.First(&cat, id)
	return &cat, nil
}

func UpdateCategoryService(id int, input *models.Category) (*models.Category, error) {
	var cat models.Category
	database.DB.First(&cat, id)
	cat.Name = input.Name
	database.DB.Save(&cat)
	return &cat, nil
}

func DeleteCategoryService(id int) error {
	return database.DB.Delete(&models.Category{}, id).Error
}
