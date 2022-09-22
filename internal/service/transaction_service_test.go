package service

import (
	"errors"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTransactionService(t *testing.T) {
	type args struct {
		c *TSConfig
	}
	tests := []struct {
		name string
		args args
		want TransactionService
	}{
		{
			name: "Test new transaction service",
			args: args{
				c: &TSConfig{
					TransactionRepository:  mocks.NewTransactionRepository(t),
					WalletRepository:       mocks.NewWalletRepository(t),
					SourceOfFundRepository: mocks.NewSourceOfFundRepository(t),
				},
			},
			want: NewTransactionService(&TSConfig{
				TransactionRepository:  mocks.NewTransactionRepository(t),
				WalletRepository:       mocks.NewWalletRepository(t),
				SourceOfFundRepository: mocks.NewSourceOfFundRepository(t),
			}),
		},
		{
			name: "Test nill transaction service",
			args: args{
				c: &TSConfig{},
			},
			want: NewTransactionService(&TSConfig{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewTransactionService(tt.args.c), "NewUserService(%v)", tt.args.c)
		})
	}
}

func Test_transactionService_GetTransactions(t *testing.T) {
	transactionRepository := mocks.NewTransactionRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	sourceOfFundRepository := mocks.NewSourceOfFundRepository(t)
	transactionService := NewTransactionService(&TSConfig{TransactionRepository: transactionRepository, WalletRepository: walletRepository, SourceOfFundRepository: sourceOfFundRepository})

	t.Run("test success get all transactions", func(t *testing.T) {
		sourceOfFundID := uint(1)
		transactionRepository.Mock.On("FindAll", 1, &dto.TransactionRequestQuery{}).
			Return([]*model.Transaction{{ID: 1, SourceOfFundID: &sourceOfFundID, SourceOfFund: &model.SourceOfFund{ID: 1, Name: "Cash"},
				UserID: 1, User: model.User{ID: 1, Name: "nabil", Email: "nabil@user.com"},
				DestinationID: 1, Wallet: model.Wallet{ID: 1, UserID: 1, User: model.User{ID: 1, Name: "nabil", Email: "nabil@user.com"}, Number: "100001", Balance: 10000},
				Amount: 50000, Description: "Top up from cash", Category: "Top Up", CreatedAt: time.Now(), UpdatedAt: time.Now()}}, nil).
			Once()

		transactions, err := transactionService.GetTransactions(1, &dto.TransactionRequestQuery{})

		assert.Nil(t, err)
		assert.Equal(t, 1, len(transactions))
	})

	t.Run("test error db when get all transactions", func(t *testing.T) {
		transactionRepository.Mock.On("FindAll", 1, &dto.TransactionRequestQuery{}).
			Return([]*model.Transaction{}, errors.New("something went wrong")).
			Once()

		transactions, err := transactionService.GetTransactions(1, &dto.TransactionRequestQuery{})

		assert.NotNil(t, err)
		assert.Equal(t, 0, len(transactions))
	})
}

func Test_transactionService_TopUp(t *testing.T) {
	transactionRepository := mocks.NewTransactionRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	sourceOfFundRepository := mocks.NewSourceOfFundRepository(t)
	transactionService := NewTransactionService(&TSConfig{TransactionRepository: transactionRepository, WalletRepository: walletRepository, SourceOfFundRepository: sourceOfFundRepository})

	t.Run("test success top up", func(t *testing.T) {
		sourceOfFundID := uint(1)
		sourceOfFundRepository.Mock.On("FindById", 1).Return(&model.SourceOfFund{ID: 1, Name: "Bank Transfer"}, nil).Once()
		walletRepository.Mock.On("FindByUserId", 1).Return(&model.Wallet{ID: 1}, nil).Once()
		transactionRepository.Mock.On("Save", mock.Anything).Return(&model.Transaction{ID: 1, SourceOfFundID: &sourceOfFundID,
			SourceOfFund: &model.SourceOfFund{ID: 1, Name: "Bank Transfer"},
			UserID:       1, User: model.User{ID: 1}, DestinationID: 1, Wallet: model.Wallet{ID: 1, Balance: 100000}, Amount: 100000,
			Description: "Top Up from Bank Transfer", Category: "Top Up"}, nil).
			Once()
		walletRepository.Mock.On("Update", &model.Wallet{ID: 1, Balance: 100000}).Return(&model.Wallet{ID: 1, Balance: 100000}, nil).Once()

		input := &dto.TopUpRequestBody{
			Amount:         100000,
			SourceOfFundID: 1,
			User:           &model.User{ID: 1},
		}
		transactions, err := transactionService.TopUp(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), transactions.ID)
	})
}

func Test_transactionService_CountTransaction(t *testing.T) {
	transactionRepository := mocks.NewTransactionRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	sourceOfFundRepository := mocks.NewSourceOfFundRepository(t)
	transactionService := NewTransactionService(&TSConfig{TransactionRepository: transactionRepository, WalletRepository: walletRepository, SourceOfFundRepository: sourceOfFundRepository})

	t.Run("test success count transaction", func(t *testing.T) {
		transactionRepository.Mock.On("Count", 1).Return(int64(28), nil).Once()

		totalTransactions, err := transactionService.CountTransaction(1)

		assert.Nil(t, err)
		assert.Equal(t, int64(28), totalTransactions)
	})

	t.Run("test error db when count transaction", func(t *testing.T) {
		transactionRepository.Mock.On("Count", 1).Return(int64(0), errors.New("something went wrong")).Once()

		totalTransactions, err := transactionService.CountTransaction(1)

		assert.NotNil(t, err)
		assert.Equal(t, int64(0), totalTransactions)
	})
}

func Test_transactionService_Transfer(t *testing.T) {
	transactionRepository := mocks.NewTransactionRepository(t)
	walletRepository := mocks.NewWalletRepository(t)
	sourceOfFundRepository := mocks.NewSourceOfFundRepository(t)
	transactionService := NewTransactionService(&TSConfig{TransactionRepository: transactionRepository, WalletRepository: walletRepository, SourceOfFundRepository: sourceOfFundRepository})

	t.Run("test success count transaction", func(t *testing.T) {
		myWallet := &model.Wallet{ID: 1, UserID: 1, Number: "1000001", Balance: 99999999}
		walletRepository.Mock.On("FindByUserId", 1).Return(myWallet, nil).Once()

		destinationWallet := &model.Wallet{ID: 2, UserID: 2, User: model.User{ID: 2, Name: "mario", Email: "mario@gmail.com"},
			Number: "100002", Balance: 1000000}
		walletRepository.Mock.On("FindByNumber", "100002").Return(destinationWallet, nil).Once()

		transactionRepository.Mock.On("Save", &model.Transaction{UserID: 2, DestinationID: 1, Amount: 500000, Description: "beli somay dek", Category: "Receive Money"}).
			Return(&model.Transaction{ID: 1, UserID: 2, DestinationID: 1, Amount: 500000, Description: "beli somay dek", Category: "Receive Money"}, nil).Once()
		transactionRepository.Mock.On("Save", &model.Transaction{UserID: 1, DestinationID: 2, Amount: 500000, Description: "beli somay dek", Category: "Send Money"}).
			Return(&model.Transaction{ID: 2, UserID: 1, DestinationID: 2, Amount: 500000, Description: "beli somay dek", Category: "Send Money"}, nil).Once()
		myWallet.Balance = myWallet.Balance - 500000
		walletRepository.Mock.On("Update", myWallet).Return(myWallet, nil).Once()
		destinationWallet.Balance = destinationWallet.Balance + 500000
		walletRepository.Mock.On("Update", destinationWallet).Return(destinationWallet, nil).Once()

		input := &dto.TransferRequestBody{}
		input.Amount = 500000
		input.Description = "beli somay dek"
		input.WalletNumber = 100002
		input.User = &model.User{ID: 1, Name: "nabil", Email: "nabil@user.com"}
		transactions, err := transactionService.Transfer(input)

		assert.Nil(t, err)
		assert.Equal(t, uint(2), transactions.ID)
		assert.Equal(t, uint(1), transactions.UserID)
		assert.Equal(t, uint(2), transactions.DestinationID)
		assert.Equal(t, 500000, transactions.Amount)
		assert.Equal(t, "Send Money", transactions.Category)
		assert.Equal(t, "beli somay dek", transactions.Description)
	})
}
