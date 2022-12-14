package repositories

import (
	"test_backend/models"

	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account models.Account) (models.Account, error)
	FindAccounts() ([]models.Account, error)
}

func RepositoryAccount(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAccount(account models.Account) (models.Account, error) {
	err := r.db.Create(&account).Error
	return account, err
}

func (r *repository) FindAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}
