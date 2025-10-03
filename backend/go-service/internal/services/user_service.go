package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
	"fmt"
)

func CreateUserService(input *models.User) (*models.User, error) {
	var existing models.User

	if err := database.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("user with email %s already exists", input.Email)
	}

	if err := database.DB.Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func GetAllUsersService() ([]models.User, error) {
	var users []models.User
	database.DB.Find(&users)
	return users, nil
}

func GetUserByIDService(id int) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserService(id int, input *models.User) (*models.User, error) {
	var user models.User
	database.DB.First(&user, id)
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	database.DB.Save(&user)
	return &user, nil
}

func DeleteUserService(id int) error {
	return database.DB.Delete(&models.User{}, id).Error
}
