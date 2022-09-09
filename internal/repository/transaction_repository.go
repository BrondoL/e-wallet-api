package repository

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error)
	Count(userID int) (int64, error)
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

func (r *transactionRepository) FindAll(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	offset := (query.Page - 1) * query.Limit
	orderBy := query.SortBy + " " + query.Sort
	queryBuider := r.db.Limit(query.Limit).Offset(offset).Order(orderBy)
	err := queryBuider.Where("user_id = ?", userID).Where("description ILIKE ?", "%"+query.Search+"%").Preload("SourceOfFund").
		Preload("User").Preload("Wallet.User").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) Count(userID int) (int64, error) {
	var total int64
	db := r.db.Model(&model.Transaction{}).Where("user_id = ?", userID).Count(&total)

	if db.Error != nil {
		return 0, db.Error
	}

	return total, nil
}

func (r *transactionRepository) Save(transaction *model.Transaction) (*model.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
