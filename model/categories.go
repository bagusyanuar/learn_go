package model

type Categories struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	Slug string
	Icon *string
}

func (category *Categories) TableName() string {
	return "categories"
}
