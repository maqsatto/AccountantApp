package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
)

func CreateAccountService(acc *models.Account) (*models.Account, error) {
	database.DB.Create(acc)
	return acc, nil
}

func GetAllAccountsService() ([]models.Account, error) {
	var accounts []models.Account
	database.DB.Find(&accounts)
	return accounts, nil
}

func GetAccountByIDService(id int) (*models.Account, error) {
	var acc models.Account
	database.DB.First(&acc, id)
	return &acc, nil
}

func UpdateAccountService(id int, input *models.Account) (*models.Account, error) {
	var acc models.Account
	database.DB.First(&acc, id)
	acc.Name = input.Name
	acc.Balance = input.Balance
	database.DB.Save(&acc)
	return &acc, nil
}

func DeleteAccountService(id int) error {
	return database.DB.Delete(&models.Account{}, id).Error
}
