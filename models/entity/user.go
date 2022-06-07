package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `gorm:"primaryKey;column:id;type:varchar(55)" json:"user_id"`
	PersonalNumber string         `gorm:"column:personal_number;type:varchar(65)" json:"personal_number"`
	Password       string         `gorm:"column:password;type:varchar(65)" json:"password"`
	Email          string         `gorm:"column:email;type:varchar(255)" json:"email"`
	Name           string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Active         bool           `gorm:"column:active;type:bool" json:"active"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
