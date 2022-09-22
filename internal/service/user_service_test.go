package service

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		c *USConfig
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		{
			name: "Test new user service",
			args: args{
				c: &USConfig{
					UserRepository:   mocks.NewUserRepository(t),
					WalletRepository: mocks.NewWalletRepository(t),
				},
			},
			want: NewUserService(&USConfig{
				UserRepository:   mocks.NewUserRepository(t),
				WalletRepository: mocks.NewWalletRepository(t),
			}),
		},
		{
			name: "Test nill user service",
			args: args{
				c: &USConfig{},
			},
			want: NewUserService(&USConfig{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewUserService(tt.args.c), "NewUserService(%v)", tt.args.c)
		})
	}
}

func Test_userService_GetUser(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	userService := NewUserService(&USConfig{UserRepository: userRepository, WalletRepository: walletRepository})

	t.Run("test success get user", func(t *testing.T) {
		userRepository.Mock.On("FindById", 1).Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@user.com", Password: "12345"}, nil).Once()

		input := &dto.UserRequestParams{}
		input.UserID = 1
		user, err := userService.GetUser(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(input.UserID), user.ID)
		assert.Equal(t, "nabil", user.Name)
		assert.Equal(t, "nabil@user.com", user.Email)
	})

	t.Run("test error db when get user", func(t *testing.T) {
		userRepository.Mock.On("FindById", -1).Return(&model.User{}, errors.New("something went wrong")).Once()

		input := &dto.UserRequestParams{}
		input.UserID = -1
		user, err := userService.GetUser(input)

		assert.NotNil(t, err)
		assert.Equal(t, uint(0), user.ID)
		assert.Equal(t, "", user.Name)
		assert.Equal(t, "", user.Email)
		assert.Equal(t, "", user.Password)
	})
}

func Test_userService_CreateUser(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	userService := NewUserService(&USConfig{UserRepository: userRepository, WalletRepository: walletRepository})

	t.Run("test success create user", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "nabil@user.com").Return(&model.User{}, nil).Once()
		userRepository.Mock.On("Save", mock.Anything).Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@user.com", Password: "$2a$04$Zw.U9pmOuDXif7bQ2hwVBOWkV3HOMDM6sdzetwtTrucoYWseqtYje"}, nil).Once()

		input := &dto.RegisterRequestBody{}
		input.Name = "nabil"
		input.Email = "nabil@user.com"
		input.Password = "12345"
		user, err := userService.CreateUser(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), user.ID)
		assert.Equal(t, "nabil", user.Name)
		assert.Equal(t, "nabil@user.com", user.Email)
	})

	t.Run("test error email not valid when create user", func(t *testing.T) {
		input := &dto.RegisterRequestBody{}
		input.Name = "nabil"
		input.Email = "nabil.com"
		input.Password = "12345"
		user, err := userService.CreateUser(input)

		assert.NotNil(t, err)
		assert.Equal(t, &custom_error.NotValidEmailError{}, err)
		assert.Equal(t, uint(0), user.ID)
		assert.Equal(t, "", user.Name)
		assert.Equal(t, "", user.Email)
	})

	t.Run("test error db find user when create user", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "nabil@user.com").Return(&model.User{}, errors.New("something went wrong")).Once()

		input := &dto.RegisterRequestBody{}
		input.Name = "nabil"
		input.Email = "nabil@user.com"
		input.Password = "12345"
		user, err := userService.CreateUser(input)

		assert.NotNil(t, err)
		assert.Equal(t, uint(0), user.ID)
		assert.Equal(t, "", user.Name)
		assert.Equal(t, "", user.Email)
	})

	t.Run("test error user already exists when create user", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "nabil@user.com").Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@user.com", Password: "12345"}, nil).Once()

		input := &dto.RegisterRequestBody{}
		input.Name = "nabil"
		input.Email = "nabil@user.com"
		input.Password = "12345"
		user, err := userService.CreateUser(input)

		assert.NotNil(t, err)
		assert.Equal(t, &custom_error.UserAlreadyExistsError{}, err)
		assert.Equal(t, uint(1), user.ID)
		assert.Equal(t, "nabil", user.Name)
		assert.Equal(t, "nabil@user.com", user.Email)
	})

	t.Run("test error db save user when create user", func(t *testing.T) {
		userRepository.Mock.On("FindByEmail", "nabil@user.com").Return(&model.User{}, nil).Once()
		userRepository.Mock.On("Save", mock.Anything).Return(&model.User{ID: 0, Name: "nabil", Email: "nabil@user.com", Password: "$2a$04$Zw.U9pmOuDXif7bQ2hwVBOWkV3HOMDM6sdzetwtTrucoYWseqtYje"}, errors.New("something went wrong")).Once()

		input := &dto.RegisterRequestBody{}
		input.Name = "nabil"
		input.Email = "nabil@user.com"
		input.Password = "12345"
		user, err := userService.CreateUser(input)

		assert.NotNil(t, err)
		assert.Equal(t, uint(0), user.ID)
		assert.Equal(t, "nabil", user.Name)
		assert.Equal(t, "nabil@user.com", user.Email)
	})
}
