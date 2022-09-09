package repository

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(userID int) ([]*model.Transaction, error)
	Save(transaction *model.Transaction) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

type TRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(c *TRConfig) TransactionRepository {
	return &transactionRepository{
		db: c.DB,
	}
}

func (r *transactionRepository) FindAll(userID int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	err := r.db.Where("user_id = ?", userID).Preload("SourceOfFund").Preload("User").Preload("Wallet.User").Order("updated_at DESC").Limit(10).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) Save(transaction *model.Transaction) (*model.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
