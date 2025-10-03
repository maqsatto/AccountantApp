package services

import (
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/models"
	"errors"
	"time"
)

func CreateTransaction(tx *models.Transaction) (*models.Transaction, error) {
	var account models.Account
	if err := database.DB.First(&account, tx.AccountID).Error; err != nil {
		return nil, errors.New("account not found")
	}

	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	if err := database.DB.Create(tx).Error; err != nil {
		return nil, err
	}

	account.Balance += tx.Amount
	if err := database.DB.Save(&account).Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func UpdateTransaction(id int, input *models.Transaction) (*models.Transaction, error) {
	var tx models.Transaction
	if err := database.DB.First(&tx, id).Error; err != nil {
		return nil, errors.New("transaction not found")
	}

	var account models.Account
	if err := database.DB.First(&account, tx.AccountID).Error; err != nil {
		return nil, errors.New("account not found")
	}

	account.Balance -= tx.Amount

	tx.Amount = input.Amount
	tx.Description = input.Description
	tx.CategoryID = input.CategoryID
	tx.AccountID = input.AccountID
	tx.Date = input.Date
	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	if err := database.DB.Save(&tx).Error; err != nil {
		return nil, err
	}

	var newAccount models.Account
	if err := database.DB.First(&newAccount, tx.AccountID).Error; err != nil {
		return nil, errors.New("account not found")
	}
	newAccount.Balance += tx.Amount
	if err := database.DB.Save(&newAccount).Error; err != nil {
		return nil, err
	}

	return &tx, nil
}

func GetAllTransactionsService() ([]models.Transaction, error) {
	var txs []models.Transaction
	if err := database.DB.Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}

func GetTransactionByIDService(id int) (*models.Transaction, error) {
	var tx models.Transaction
	if err := database.DB.First(&tx, id).Error; err != nil {
		return nil, errors.New("transaction not found")
	}
	return &tx, nil
}

func DeleteTransaction(id int) error {
	var tx models.Transaction
	if err := database.DB.First(&tx, id).Error; err != nil {
		return errors.New("transaction not found")
	}

	var account models.Account
	if err := database.DB.First(&account, tx.AccountID).Error; err != nil {
		return errors.New("account not found")
	}

	account.Balance -= tx.Amount
	if err := database.DB.Save(&account).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&tx).Error; err != nil {
		return err
	}

	return nil
}
