package repository

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type SourceOfFundRepository interface {
	FindById(id int) (*model.SourceOfFund, error)
}

type sourceOfFundRepository struct {
	db *gorm.DB
}

type SRConfig struct {
	DB *gorm.DB
}

func NewSourceOfFundRepository(c *SRConfig) SourceOfFundRepository {
	return &sourceOfFundRepository{
		db: c.DB,
	}
}

func (r *sourceOfFundRepository) FindById(id int) (*model.SourceOfFund, error) {
	var sourceOfFund *model.SourceOfFund

	err := r.db.Where("id = ?", id).Find(&sourceOfFund).Error
	if err != nil {
		return sourceOfFund, err
	}

	return sourceOfFund, nil
}
