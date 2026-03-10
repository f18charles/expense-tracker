package models

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Name      string    `gorm:"not null" json:"name"`
	Type      string    `gorm:"not null" json:"type"`
	Color     string    `json:"color"`
	Icon      string    `json:"icon"`
	IsDefault bool      `gorm:"default:false" json:"is_default"`
	CreatedAt time.Time `json:"created_at"`

	User *User `gorm:"foreignKey:UserID" json:"-"`
}
