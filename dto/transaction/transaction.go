package transactiondto

type RequestTransaction struct {
	ID                int    `json:"id" `
	AccountID         int    `json:"account_id" validate:"required"`
	Desc              string `json:"desc" validate:"required"`
	DebitCreditStatus string `json:"status" validate:"required"`
	Amount            int    `json:"amount" validate:"required"`
}

type GetTransactions struct {
	AccountID int `json:"account_id" validate:"required"`
}
