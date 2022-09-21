package dto

import (
	"strings"
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
	WalletNumber int    `json:"wallet_number" binding:"required"`
	Description  string `json:"description"`
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

type TransactionRequestQuery struct {
	Search string `form:"s"`
	SortBy string `form:"sortBy"`
	Sort   string `form:"sort"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type Destination struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type TransferResponse struct {
	ID            uint        `json:"id"`
	Destination   Destination `json:"destination"`
	Amount        int         `json:"amount"`
	WalletBalance int         `json:"balance"`
	Description   string      `json:"description"`
	Category      string      `json:"category"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type TransactionResponse struct {
	ID           uint        `json:"id"`
	SourceOfFund string      `json:"source_of_fund"`
	Destination  Destination `json:"destination"`
	Amount       int         `json:"amount"`
	Description  string      `json:"description"`
	Category     string      `json:"category"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
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

func FormatTransfer(transaction *model.Transaction) TransferResponse {
	return TransferResponse{
		ID:            transaction.ID,
		Destination:   Destination{Name: transaction.Wallet.User.Name, Number: transaction.Wallet.Number},
		Amount:        transaction.Amount,
		WalletBalance: int(*transaction.SourceOfFundID),
		Description:   transaction.Description,
		Category:      transaction.Category,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func FormatTransaction(transaction *model.Transaction) TransactionResponse {
	var sourceOfFund string
	if transaction.SourceOfFund != nil {
		sourceOfFund = transaction.SourceOfFund.Name
	}
	return TransactionResponse{
		ID:           transaction.ID,
		SourceOfFund: sourceOfFund,
		Destination:  Destination{Name: transaction.Wallet.User.Name, Number: transaction.Wallet.Number},
		Amount:       transaction.Amount,
		Description:  transaction.Description,
		Category:     transaction.Category,
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,
	}
}

func FormatTransactions(transactions []*model.Transaction) []TransactionResponse {
	formattedTransactions := []TransactionResponse{}
	for _, transaction := range transactions {
		formattedBook := FormatTransaction(transaction)
		formattedTransactions = append(formattedTransactions, formattedBook)
	}
	return formattedTransactions
}

func FormatQuery(query *TransactionRequestQuery) *TransactionRequestQuery {
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	query.SortBy = strings.ToLower(query.SortBy)
	if query.SortBy == "date" {
		query.SortBy = "updated_at"
	} else if query.SortBy == "to" {
		query.SortBy = "destination_id"
	} else if query.SortBy != "amount" {
		query.SortBy = "updated_at"
	}

	query.Sort = strings.ToUpper(query.Sort)
	if query.Sort != "ASC" {
		query.Sort = "DESC"
	}

	return query
}
