package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	PassHash string
}

type Transaction struct {
	gorm.Model
	Date        time.Time
	Description string
	Category    string
	Amount      int64
	IsIncome    bool
	UserID      uint
	User        User
}

