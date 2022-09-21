package service

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletService(t *testing.T) {
	type args struct {
		c *WSConfig
	}
	tests := []struct {
		name string
		args args
		want WalletService
	}{
		{
			name: "Test new wallet service",
			args: args{
				c: &WSConfig{
					UserRepository:   mocks.NewUserRepository(t),
					WalletRepository: mocks.NewWalletRepository(t),
				},
			},
			want: NewWalletService(&WSConfig{
				UserRepository:   mocks.NewUserRepository(t),
				WalletRepository: mocks.NewWalletRepository(t),
			}),
		},
		{
			name: "Test nill wallet service",
			args: args{
				c: &WSConfig{},
			},
			want: NewWalletService(&WSConfig{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewWalletService(tt.args.c), "NewWalletService(%v)", tt.args.c)
		})
	}
}

func Test_walletService_GetWalletByUserId(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	walletService := NewWalletService(&WSConfig{UserRepository: userRepository, WalletRepository: walletRepository})

	t.Run("test success get wallet by user id", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = 1
		walletRepository.Mock.On("FindByUserId", input.UserID).Return(&model.Wallet{ID: 1, UserID: uint(input.UserID), Number: "100001", Balance: 1000000}, nil).Once()

		wallet, err := walletService.GetWalletByUserId(input)
		assert.Nil(t, err)
		assert.Equal(t, uint(input.UserID), wallet.UserID)
	})

	t.Run("test success get wallet by user id", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = -1
		walletRepository.Mock.On("FindByUserId", input.UserID).Return(&model.Wallet{}, errors.New("something went wrong")).Once()

		wallet, err := walletService.GetWalletByUserId(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), wallet.UserID)
	})
}

func Test_walletService_CreateWallet(t *testing.T) {
	userRepository := mocks.NewUserRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	walletService := NewWalletService(&WSConfig{UserRepository: userRepository, WalletRepository: walletRepository})

	t.Run("test success create wallet", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = 1

		userRepository.Mock.On("FindById", input.UserID).Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@gmail.com", Password: "12345"}, nil).Once()
		walletRepository.Mock.On("FindByUserId", input.UserID).Return(&model.Wallet{}, nil).Once()
		wallet := &model.Wallet{UserID: uint(input.UserID), Number: "100001", Balance: 0}
		walletRepository.Mock.On("Save", wallet).Return(&model.Wallet{ID: 1, UserID: uint(input.UserID), Number: "100001", Balance: 0}, nil).Once()

		wallet, err := walletService.CreateWallet(input)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), wallet.ID)
		assert.Equal(t, uint(1), wallet.UserID)
		assert.Equal(t, "100001", wallet.Number)
		assert.Equal(t, 0, wallet.Balance)
	})

	t.Run("test error db find user when create wallet", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = -1
		userRepository.Mock.On("FindById", input.UserID).Return(&model.User{}, errors.New("something went wrong")).Once()

		wallet, err := walletService.CreateWallet(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), wallet.ID)
		assert.Equal(t, uint(0), wallet.UserID)
		assert.Equal(t, "", wallet.Number)
		assert.Equal(t, 0, wallet.Balance)
	})

	t.Run("test error user not found when create wallet", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = -1
		userRepository.Mock.On("FindById", input.UserID).Return(&model.User{}, nil).Once()

		wallet, err := walletService.CreateWallet(input)
		assert.NotNil(t, err)
		assert.Equal(t, &custom_error.UserNotFoundError{}, err)
		assert.Equal(t, uint(0), wallet.ID)
		assert.Equal(t, uint(0), wallet.UserID)
		assert.Equal(t, "", wallet.Number)
		assert.Equal(t, 0, wallet.Balance)
	})

	t.Run("test error db find wallet when create wallet", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = 1
		userRepository.Mock.On("FindById", input.UserID).Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@gmail.com", Password: "12345"}, nil).Once()
		walletRepository.Mock.On("FindByUserId", input.UserID).Return(&model.Wallet{}, errors.New("something went wrong")).Once()

		wallet, err := walletService.CreateWallet(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), wallet.ID)
		assert.Equal(t, uint(0), wallet.UserID)
		assert.Equal(t, "", wallet.Number)
		assert.Equal(t, 0, wallet.Balance)
	})

	t.Run("test error wallet already exists when create wallet", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = 1
		userRepository.Mock.On("FindById", input.UserID).Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@gmail.com", Password: "12345"}, nil).Once()
		walletRepository.Mock.On("FindByUserId", input.UserID).Return(&model.Wallet{ID: 1, UserID: uint(input.UserID), Number: "100001", Balance: 0}, nil).Once()

		wallet, err := walletService.CreateWallet(input)
		assert.NotNil(t, err)
		assert.Equal(t, &custom_error.WalletAlreadyExistsError{}, err)
		assert.Equal(t, uint(0), wallet.ID)
		assert.Equal(t, uint(0), wallet.UserID)
		assert.Equal(t, "", wallet.Number)
		assert.Equal(t, 0, wallet.Balance)
	})

	t.Run("test error wallet already exists when create wallet", func(t *testing.T) {
		input := &dto.WalletRequestBody{}
		input.UserID = 1
		userRepository.Mock.On("FindById", input.UserID).Return(&model.User{ID: 1, Name: "nabil", Email: "nabil@gmail.com", Password: "12345"}, nil).Once()
		walletRepository.Mock.On("FindByUserId", input.UserID).Return(&model.Wallet{}, nil).Once()
		wallet := &model.Wallet{UserID: uint(input.UserID), Number: "100001", Balance: 0}
		walletRepository.Mock.On("Save", wallet).Return(&model.Wallet{}, errors.New("something went wrong")).Once()

		wallet, err := walletService.CreateWallet(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), wallet.ID)
		assert.Equal(t, uint(0), wallet.UserID)
		assert.Equal(t, "", wallet.Number)
		assert.Equal(t, 0, wallet.Balance)
	})
}
