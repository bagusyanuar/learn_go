package model

type Subject struct {
	ID   uint   `gorm:"primaryKey;type:int(11);" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Slug string `gorm:"column:slug" json:"slug"`
	Icon string `gorm:"column:icon;type:text" json:"icon"`
}

func (pc *Subject) TableName() string {
	return "subjects"
}
