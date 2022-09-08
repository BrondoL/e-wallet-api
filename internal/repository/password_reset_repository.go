package repository

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type PassowrdResetRepository interface {
	FindByEmail(email string) (*model.PasswordReset, error)
	Save(passwordReset *model.PasswordReset) (*model.PasswordReset, error)
}

type passwordResetRepository struct {
	db *gorm.DB
}

type PRConfig struct {
	DB *gorm.DB
}

func NewPasswordResetRepository(c *PRConfig) PassowrdResetRepository {
	return &passwordResetRepository{
		db: c.DB,
	}
}

func (r *passwordResetRepository) FindByEmail(email string) (*model.PasswordReset, error) {
	var passwordReset *model.PasswordReset

	err := r.db.Where("email = ?", email).Find(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (r *passwordResetRepository) Save(passwordReset *model.PasswordReset) (*model.PasswordReset, error) {
	err := r.db.Save(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}