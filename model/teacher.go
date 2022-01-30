package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//abstract for model Teacher
type Teacher struct {
	Id        uuid.UUID `gorm:"primaryKey;type:char(36);" json:"id"`
	FullName  string    `gorm:"column:full_name"`
	Slug      string    `gorm:"column:slug" json:"slug"`
	Grade     uint8     `gorm:"column:grade;type:smallint" json:"grade"`
	Gender    uint8     `gorm:"column:gender" json:"gender"`
	Address   string    `gorm:"column:address" json:"address"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UserId    uuid.UUID `gorm:"column:user_id"`
}

type TeachersApi struct {
	Teacher
	Alias string `gorm:"->" json:"alias"`
}

func (teacher *Teacher) TableName() string {
	return "teacher"
}

func (teacher *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	teacher.Id = uuid.New()
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()
	return
}
