package service

import (
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAuthService(t *testing.T) {
	var authService = NewAuthService(&ASConfig{UserRepository: mocks.NewUserRepository(t), PasswordResetRepository: mocks.NewPassowrdResetRepository(t)})

	assert.NotNil(t, authService)
}

func Test_authService_Attempt(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	authService := NewAuthService(&ASConfig{UserRepository: userRepository, PasswordResetRepository: mocks.NewPassowrdResetRepository(t)})

	t.Run("test success attempt user", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "nabil@user.com").Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@user.com", Password: "$2a$04$93AZUXoqhOu6TNb481MYke3iDbM8UAzizOHmKSEf36bQtzV3kffwm"}, nil).Once()

		input := &dto.LoginRequestBody{}
		input.Email = "nabil@user.com"
		input.Password = "12345"
		user, err := authService.Attempt(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), user.ID)
		assert.Equal(t, "nabil", user.Name)
		assert.Equal(t, "nabil@user.com", user.Email)
	})
}

func Test_authService_ForgotPass(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	passwordResetRepository := mocks.NewPassowrdResetRepository(t)
	authService := NewAuthService(&ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})

	t.Run("test success forgot password", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "nabil@user.com").Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@user.com", Password: "$2a$04$93AZUXoqhOu6TNb481MYke3iDbM8UAzizOHmKSEf36bQtzV3kffwm"}, nil).Once()
		passwordResetRepository.Mock.On("FindByUserId", 1).Return(&model.PasswordReset{}, nil).Once()
		passwordResetRepository.Mock.On("Save", mock.Anything).Return(&model.PasswordReset{ID: 1, UserID: 1, Token: "brondol", ExpiredAt: time.Now().Add(time.Minute * 15)}, nil).Once()

		input := &dto.ForgotPasswordRequestBody{}
		input.Email = "nabil@user.com"
		passwordReset, err := authService.ForgotPass(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), passwordReset.ID)
		assert.Equal(t, "brondol", passwordReset.Token)
	})
}

func Test_authService_ResetPass(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	passwordResetRepository := mocks.NewPassowrdResetRepository(t)
	authService := NewAuthService(&ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})

	t.Run("test success reset password", func(t *testing.T) {
		user := model.User{ID: 1, Name: "nabil", Email: "nabil@user.com", Password: "12345"}
		passwordReset := &model.PasswordReset{ID: 1, UserID: 1, User: user, Token: "brondol", ExpiredAt: time.Now().Add(time.Minute * 15)}
		passwordResetRepository.Mock.On("FindByToken", "brondol").Return(passwordReset, nil).Once()
		userRepository.Mock.On("Update", mock.Anything).Return(&user, nil).Once()
		passwordResetRepository.Mock.On("Delete", passwordReset).Return(passwordReset, nil).Once()

		input := &dto.ResetPasswordRequestBody{}
		input.Token = "brondol"
		input.Password = "12345"
		input.ConfirmPassword = "12345"
		passwordReset, err := authService.ResetPass(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), passwordReset.ID)
		assert.Equal(t, "brondol", passwordReset.Token)
	})
}
