package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

//abstract User
type User struct {
	Id         uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	IsVerified bool           `gorm:"column:is_verified" json:"isVerified"`
	Roles      datatypes.JSON `json:"roles"`
	Provider   datatypes.JSON `json:"provider"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
}

func (user *User) TableName() string {
	return "users"

}
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return
}

//User Associate with Member
type UserMember struct {
	User
	Member Member `gorm:"foreignKey:UserId" json:"member"`
}

type UserTeacher struct {
	User
	Teacher Teacher `gorm:"foreignKey:UserId" json:"teacher"`
}
