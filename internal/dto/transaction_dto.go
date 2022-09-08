package dto

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
)

type TopUpRequestBody struct {
	Amount         int `json:"amount" binding:"required,min=50000,max=10000000"`
	SourceOfFundID int `json:"source_of_fund_id" binding:"required"`
	User           *model.User
}

type TransferRequestBody struct {
	Amount       int    `json:"amount" binding:"required,min=1000,max=50000000"`
	WalletNumber string `json:"wallet_number" binding:"required"`
	User         *model.User
}

type TopUpResponse struct {
	ID            uint      `json:"id"`
	SourceOfFund  string    `json:"source_of_fund"`
	Amount        int       `json:"amount"`
	WalletBalance int       `json:"balance"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func FormatTopUp(transaction *model.Transaction) TopUpResponse {
	return TopUpResponse{
		ID:            transaction.ID,
		SourceOfFund:  transaction.SourceOfFund.Name,
		Amount:        transaction.Amount,
		WalletBalance: transaction.Wallet.Balance,
		Description:   transaction.Description,
		Category:      transaction.Category,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}
