package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID              uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID          uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	AccountID       uuid.UUID  `gorm:"type:uuid;not null" json:"account_id"`
	CategoryID      *uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Amount          float64    `gorm:"type:numeric(15,2);not null" json:"amount"`
	Type            string     `gorm:"not null" json:"type"`
	Description     string     `json:"description"`
	PaymentMethod   string     `json:"payment_method"`
	ReferenceID     string     `json:"reference_id"`
	Status          string     `gorm:"default:completed" json:"status"`
	TransactionDate time.Time  `json:"transaction_date"`
	CreatedAt       time.Time  `json:"created_at"`

	User     User      `gorm:"foreignKey:UserID" json:"-"`
	Account  Account   `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}
