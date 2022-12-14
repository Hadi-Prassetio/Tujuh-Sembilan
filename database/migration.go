package database

import (
	"fmt"
	"test_backend/models"
	"test_backend/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(models.Account{}, models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("migration error")
	}
	fmt.Println("migration success")
}
