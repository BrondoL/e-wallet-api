package dto

import "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatLogin(user *model.User, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
