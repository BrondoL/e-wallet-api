package service

import (
	"strconv"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	r "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/repository"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
)

type TransactionService interface {
	GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error)
	TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error)
	Transfer(input *dto.TransferRequestBody) (*model.Transaction, error)
	CountTransaction(userID int) (int64, error)
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

func (s *transactionService) GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll(userID, query)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
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
	transaction.SourceOfFundID = &sourceOfFund.ID
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

	transaction.SourceOfFund = sourceOfFund
	transaction.User = *input.User
	transaction.Wallet = *wallet

	return transaction, nil
}

func (s *transactionService) CountTransaction(userID int) (int64, error) {
	totalTransactions, err := s.transactionRepository.Count(userID)
	if err != nil {
		return totalTransactions, err
	}

	return totalTransactions, nil
}

func (s *transactionService) Transfer(input *dto.TransferRequestBody) (*model.Transaction, error) {
	myWallet, err := s.walletRepository.FindByUserId(int(input.User.ID))
	if err != nil {
		return &model.Transaction{}, err
	}
	if myWallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}
	if myWallet.Balance < input.Amount {
		return &model.Transaction{}, &custom_error.InsufficientBallanceError{}
	}
	number := strconv.Itoa(input.WalletNumber)
	if myWallet.Number == number {
		return &model.Transaction{}, &custom_error.TransferToSameWalletError{}
	}

	destinationWallet, err := s.walletRepository.FindByNumber(number)
	if err != nil {
		return &model.Transaction{}, err
	}
	if destinationWallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}

	//create transaction for receiver
	transaction := &model.Transaction{}
	transaction.UserID = destinationWallet.User.ID
	transaction.DestinationID = myWallet.ID
	transaction.Amount = input.Amount
	transaction.Description = input.Description
	transaction.Category = "Receive Money"

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	// create transaction for sender
	transaction = &model.Transaction{}
	transaction.UserID = input.User.ID
	transaction.DestinationID = destinationWallet.ID
	transaction.Amount = input.Amount
	transaction.Description = input.Description
	transaction.Category = "Send Money"

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	myWallet.Balance = myWallet.Balance - input.Amount
	myWallet, err = s.walletRepository.Update(myWallet)
	if err != nil {
		return transaction, err
	}

	destinationWallet.Balance = destinationWallet.Balance + input.Amount
	_, err = s.walletRepository.Update(destinationWallet)
	if err != nil {
		return transaction, err
	}

	balance := uint(myWallet.Balance)
	transaction.SourceOfFundID = &balance
	transaction.User = *input.User
	transaction.Wallet = *destinationWallet

	return transaction, nil
}
