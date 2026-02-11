package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	PassHash string
}

type Transaction struct {
	gorm.Model
	UserID      uint
	User        User
	Date        uint
	Description string
	Category    string
	Amount      int64
	IsIncome    bool
}
