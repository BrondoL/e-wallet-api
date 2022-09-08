package dto

type WalletRequestBody struct {
	UserID int `json:"name" binding:"required,alphanum"`
}
