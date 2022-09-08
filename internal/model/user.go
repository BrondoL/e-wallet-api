package model

type User struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Email         string        `gorm:"uniqueIndex"`
	PasswordReset PasswordReset `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Password      string
}

func (User) TableName() string {
	return "users"
}
