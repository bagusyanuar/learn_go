package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//abstract Member
type Member struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	FullName  string    `gorm:"column:full_name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	School    string    `gorm:"column:school" json:"school"`
	Address   string    `gorm:"column:address"`
	Avatar    string    `gorm:"column:avatar"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UserId    uuid.UUID `gorm:"column:user_id"`
}

func (member *Member) TableName() string {
	return "member"
}

func (member *Member) BeforeCreate(tx *gorm.DB) (err error) {
	member.Id = uuid.New()
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()
	return
}

//Member associate with User
type MemberUser struct {
	Member
	User User `gorm:"foreignKey:UserId" json:"user"`
}

func (member *MemberUser) TableName() string {
	return "member"
}

type ResponseMemberUser struct {
	Id       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	FullName string    `gorm:"column:full_name"`
	Phone    *string   `gorm:"column:phone" json:"phone"`
	School   *string   `gorm:"column:school" json:"school"`
	// Address   *string   `gorm:"column:address"`
	// Avatar    *string   `gorm:"column:avatar"`
	// CreatedAt time.Time `gorm:"column:created_at"`
	// UpdatedAt time.Time `gorm:"column:updated_at"`
	UserId uuid.UUID `gorm:"column:user_id" json:"user_id"`
	User   *User     `gorm:"foreignKey:user_id" json:"user"`
}

func (member *ResponseMemberUser) TableName() string {
	return "member"
}
