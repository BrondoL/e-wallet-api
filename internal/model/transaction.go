package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint `gorm:"primaryKey"`
	SourceOfFundID *uint
	SourceOfFund   *SourceOfFund `gorm:"foreignKey:SourceOfFundID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint
	User           User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DestinationID  uint
	Wallet         Wallet `gorm:"foreignKey:DestinationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount         int
	Description    string
	Category       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Transaction) TableName() string {
	return "transactions"
}
