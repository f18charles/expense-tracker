package models

import (
	"time"
)

type User struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	PassHash      string    `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	authenticated bool
}

type Transaction struct {
	ID          uint
	UserID      uint
	Date        uint
	Description string
	Category    string
	Amount      int64
	IsIncome    bool
}

type MonthlySummary struct {
	Month      time.Month
	Year       int
	OpeningBal int64
	TotalExp   int64
	CurrBal    int64
	ByCategory map[string]int64
}
