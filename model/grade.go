package model

import "time"

type Grade struct {
	ID        uint      `gorm:"primaryKey;" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Slug      string    `gorm:"column:slug;type:varchar(255)" json:"slug"`
	Icon      *string   `gorm:"column:icon;type:varchar(255)" json:"icon"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (grade *Grade) TableName() string {
	return "grades"
}
