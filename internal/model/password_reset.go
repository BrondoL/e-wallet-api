package model

import "time"

type PasswordReset struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token     string
	ExpiredAt time.Time
}

func (PasswordReset) TableName() string {
	return "password_resets"
}
