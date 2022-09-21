package dto

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
)

type UserRequestParams struct {
	UserID int `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type UserResponseBody struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user *model.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	return formattedUser
}

func FormatUsers(authors []*model.User) []UserResponseBody {
	formattedUsers := []UserResponseBody{}
	for _, user := range authors {
		formattedUser := FormatUser(user)
		formattedUsers = append(formattedUsers, formattedUser)
	}
	return formattedUsers
}

type UserDetailResponse struct {
	ID     uint           `json:"id"`
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Wallet WalletResponse `json:"wallet"`
}

func FormatUserDetail(user *model.User, wallet *model.Wallet) UserDetailResponse {
	formattedUser := UserDetailResponse{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	formattedUser.Wallet = FormatWallet(wallet)
	return formattedUser
}
