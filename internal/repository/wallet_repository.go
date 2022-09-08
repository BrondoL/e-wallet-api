package repository

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	FindByUserId(id int) (*model.Wallet, error)
	FindByNumber(number string) (*model.Wallet, error)
	Save(wallet *model.Wallet) (*model.Wallet, error)
	Update(wallet *model.Wallet) (*model.Wallet, error)
}

type walletRepository struct {
	db *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(c *WRConfig) WalletRepository {
	return &walletRepository{
		db: c.DB,
	}
}

func (r *walletRepository) FindByUserId(id int) (*model.Wallet, error) {
	var wallet *model.Wallet

	err := r.db.Where("user_id = ?", id).Find(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (r *walletRepository) FindByNumber(number string) (*model.Wallet, error) {
	var wallet *model.Wallet

	err := r.db.Where("number = ?", number).Preload("User").Find(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (r *walletRepository) Save(wallet *model.Wallet) (*model.Wallet, error) {
	err := r.db.Create(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (r *walletRepository) Update(wallet *model.Wallet) (*model.Wallet, error) {
	err := r.db.Save(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}
