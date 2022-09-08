package service

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	r "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/repository"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
)

type TransactionService interface {
	TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error)
}

type transactionService struct {
	transactionRepository  r.TransactionRepository
	walletRepository       r.WalletRepository
	sourceOfFundRepository r.SourceOfFundRepository
}

type TSConfig struct {
	TransactionRepository  r.TransactionRepository
	WalletRepository       r.WalletRepository
	SourceOfFundRepository r.SourceOfFundRepository
}

func NewTransactionService(c *TSConfig) TransactionService {
	return &transactionService{
		transactionRepository:  c.TransactionRepository,
		walletRepository:       c.WalletRepository,
		sourceOfFundRepository: c.SourceOfFundRepository,
	}
}

func (s *transactionService) TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error) {
	sourceOfFund, err := s.sourceOfFundRepository.FindById(input.SourceOfFundID)
	if err != nil {
		return &model.Transaction{}, err
	}
	if sourceOfFund.ID == 0 {
		return &model.Transaction{}, &custom_error.SourceOfFundNotFoundError{}
	}

	wallet, err := s.walletRepository.FindByUserId(int(input.User.ID))
	if err != nil {
		return &model.Transaction{}, err
	}
	if wallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}

	transaction := &model.Transaction{}
	transaction.SourceOfFundID = sourceOfFund.ID
	transaction.UserID = input.User.ID
	transaction.DestinationID = wallet.ID
	transaction.Amount = input.Amount
	transaction.Description = "Top Up from " + sourceOfFund.Name
	transaction.Category = "Top Up"

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	wallet.Balance = wallet.Balance + input.Amount
	wallet, err = s.walletRepository.Update(wallet)
	if err != nil {
		return transaction, err
	}

	transaction.SourceOfFund = *sourceOfFund
	transaction.User = *input.User
	transaction.Wallet = *wallet
	return transaction, nil
}
