package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCourse struct {
	ID         uuid.UUID `gorm:"primaryKey;type:char(36);" json:"id"`
	TeacherId  uuid.UUID `gorm:"column:teacher_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null;"`
	SubjectId  uint      `gorm:"column:subject_id;type:int(11);not null;"`
	CategoryId uint      `gorm:"column:category_id;type:int(11);not null;"`
	Slug       string    `gorm:"column:slug;type:varchar(255)" json:"slug"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

type ProductCourseRelationship struct {
	ProductCourse
	Teacher    Teacher    `gorm:"foreignKey:TeacherId;" json:"teacher"`
	Subject    Subject    `gorm:"foreignKey:SubjectId" json:"subject"`
	Categories Categories `gorm:"foreignKey:CategoryId" json:"category"`
}

func (pc *ProductCourse) TableName() string {
	return "product_course"
}
