package model

type Wallet struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	User    User
	Number  string
	Balance int
}

func (Wallet) TableName() string {
	return "wallets"
}
