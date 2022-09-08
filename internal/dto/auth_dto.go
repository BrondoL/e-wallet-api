package dto

import "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type RegisterRequestBody struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestBody struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=5"`
}

type ForgotPasswordResponseBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponseBody struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	WalletNumber string `json:"wallet"`
	Token        string `json:"token"`
}

func FormatLogin(user *model.User, wallet *model.Wallet, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		WalletNumber: wallet.Number,
		Token:        token,
	}
}

func FormatForgotPassword(passwordReset *model.PasswordReset) ForgotPasswordResponseBody {
	return ForgotPasswordResponseBody{
		Email: passwordReset.User.Email,
		Token: passwordReset.Token,
	}
}
