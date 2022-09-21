package dto

import "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"

type WalletRequestBody struct {
	UserID int `json:"name" binding:"required,alphanum"`
}

type WalletResponse struct {
	ID      int    `json:"id"`
	Number  string `json:"number"`
	Balance int    `json:"balance"`
}

func FormatWallet(wallet *model.Wallet) WalletResponse {
	return WalletResponse{
		ID:      int(wallet.ID),
		Number:  wallet.Number,
		Balance: wallet.Balance,
	}
}
