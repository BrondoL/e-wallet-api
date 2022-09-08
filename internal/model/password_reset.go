package model

import "time"

type PasswordReset struct {
	Email     string `gorm:"primaryKey"`
	Token     string
	ExpiredAt time.Time
}

func (PasswordReset) TableName() string {
	return "password_resets"
}
