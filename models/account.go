package models

import "time"

type Account struct {
	ID           int           `json:"account_id"`
	Name         string        `json:"name"`
	Password     string        `json:"-"`
	Transactions []Transaction `json:"-"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAd    time.Time     `json:"-"`
}
