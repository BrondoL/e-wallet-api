package service

import (
	"net/mail"
	"time"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	r "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/repository"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Attempt(input *dto.LoginRequestBody) (*model.User, error)
	ForgotPass(input *dto.ForgotPasswordRequestBody) (*model.PasswordReset, error)
	ResetPass(input *dto.ResetPasswordRequestBody) (*model.PasswordReset, error)
}

type authService struct {
	userRepository          r.UserRepository
	passwordResetRepository r.PassowrdResetRepository
}

type ASConfig struct {
	UserRepository          r.UserRepository
	PasswordResetRepository r.PassowrdResetRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository:          c.UserRepository,
		passwordResetRepository: c.PasswordResetRepository,
	}
}

func (s *authService) Attempt(input *dto.LoginRequestBody) (*model.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &model.User{}, &custom_error.NotValidEmailError{}
	}

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, &custom_error.UserNotFoundError{}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, &custom_error.IncorrectCredentialsError{}
	}

	return user, nil
}

func (s *authService) ForgotPass(input *dto.ForgotPasswordRequestBody) (*model.PasswordReset, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &model.PasswordReset{}, &custom_error.NotValidEmailError{}
	}

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return &model.PasswordReset{}, err
	}

	if user.ID == 0 {
		return &model.PasswordReset{}, &custom_error.UserNotFoundError{}
	}

	passwordReset, err := s.passwordResetRepository.FindByUserId(int(user.ID))
	if err != nil {
		return &model.PasswordReset{}, err
	}

	passwordReset.UserID = user.ID
	passwordReset.Token = utils.GenerateString(10)
	passwordReset.ExpiredAt = time.Now().Add(time.Minute * 15)

	passwordReset, err = s.passwordResetRepository.Save(passwordReset)
	passwordReset.User = *user

	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (s *authService) ResetPass(input *dto.ResetPasswordRequestBody) (*model.PasswordReset, error) {
	passwordReset, err := s.passwordResetRepository.FindByToken(input.Token)
	if err != nil {
		return passwordReset, err
	}

	if passwordReset.User.Email == "" {
		return passwordReset, &custom_error.ResetTokenNotFound{}
	}

	if input.Password != input.ConfirmPassword {
		return passwordReset, &custom_error.PasswordNotSame{}
	}

	user := &passwordReset.User
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return passwordReset, err
	}
	user.Password = string(passwordHash)

	_, err = s.userRepository.Update(user)
	if err != nil {
		return passwordReset, err
	}

	passwordReset, err = s.passwordResetRepository.Delete(passwordReset)
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}
