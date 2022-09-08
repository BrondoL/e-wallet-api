package dto

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
)

type UserRequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequestParams struct {
	UserID int `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name string `form:"name"`
}

type UserResponseBody struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
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

func FormatLogin(user *model.User, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
