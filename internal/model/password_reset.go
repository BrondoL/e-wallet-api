package model

import "time"

type PasswordReset struct {
	Email     string `gorm:"primaryKey"`
	Token     string
	User      User `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExpiredAt time.Time
}

func (PasswordReset) TableName() string {
	return "password_resets"
}
