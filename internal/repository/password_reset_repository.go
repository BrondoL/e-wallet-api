package repository

import (
	"time"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"gorm.io/gorm"
)

type PassowrdResetRepository interface {
	FindByUserId(id int) (*model.PasswordReset, error)
	FindByToken(token string) (*model.PasswordReset, error)
	Save(passwordReset *model.PasswordReset) (*model.PasswordReset, error)
	Delete(passwordReset *model.PasswordReset) (*model.PasswordReset, error)
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

func (r *passwordResetRepository) FindByUserId(id int) (*model.PasswordReset, error) {
	var passwordReset *model.PasswordReset

	err := r.db.Where("user_id = ?", id).Find(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (r *passwordResetRepository) FindByToken(token string) (*model.PasswordReset, error) {
	var passwordReset *model.PasswordReset

	err := r.db.Where("token = ?", token).Where("expired_at >= ?", time.Now()).Preload("User").Find(&passwordReset).Error
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

func (r *passwordResetRepository) Delete(passwordReset *model.PasswordReset) (*model.PasswordReset, error) {
	err := r.db.Delete(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}
