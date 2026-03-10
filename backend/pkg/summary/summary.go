package summary

import (
	"github.com/google/uuid"
	"time"
)

type Period string

const (
	PeriodDaily   Period = "daily"
	PeriodWeekly  Period = "weekly"
	PeriodMonthly Period = "monthly"
	PeriodYearly  Period = "yearly"
)

type MonthlySummary struct {
	UserID        uuid.UUID                `json:"-"`
	Year          int                      `json:"year"`
	Month         time.Month               `json:"month"`
	Income        float64                  `json:"income"`
	Expenses      float64                  `json:"expenses"`
	Savings       float64                  `json:"savings"`
	SavingsRate   float64                  `json:"savings_rate"`
	ByCategory    map[string]CategorySpend `json:"by_category"`
	PreviousMonth *MonthlySummary          `json:"previous_month,omitempty"`
}

type CategorySpend struct {
	CategoryID    uuid.UUID `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	CategoryColor string    `json:"color"`
	Spent         float64   `json:"spent"`
	Budget        float64   `json:"budget,omitempty"`
	Percentage    float64   `json:"percentage"`
	Trend         float64   `json:"trend"`
}
