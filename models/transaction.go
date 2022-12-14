package models

import "time"

type Transaction struct {
	ID                int       `json:"id"`
	AccountID         int       `json:"account_id"`
	Desc              string    `json:"desc"`
	DebitCreditStatus string    `json:"status"`
	Amount            int       `json:"amount"`
	CreatedAt         time.Time `json:"transaction_date"`
	UpdatedAt         time.Time `json:"-"`
}
