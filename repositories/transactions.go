package repositories

import (
	"test_backend/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	FindTransactions() ([]models.Transaction, error)
	GetAccountTransactions(ID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *repository) GetAccountTransactions(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions, "account_id = ?", ID).Error

	return transactions, err
}
