package model

type SourceOfFund struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (SourceOfFund) TableName() string {
	return "source_of_funds"
}
