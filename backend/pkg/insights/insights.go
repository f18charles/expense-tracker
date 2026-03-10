package insights

import (
	"time"

	"github.com/google/uuid"
)

type SpendingInsights struct {
	TopCategories    []CategoryTrend      `json:"top_categories"`
	SpendingPatterns SpendingPatterns     `json:"spending_patterns"`
	Anomalies        []TransactionAnomaly `json:"anomalies"`
	Recommendations  []string             `json:"recommendations"`
}

type CategoryTrend struct {
	CategoryName     string  `json:"category_name"`
	TotalSpent       float64 `json:"total_spent"`
	Percentage       float64 `json:"percentage"`
	TransactionCount int     `json:"transaction_count"`
	AvgPerTx         float64 `json:"avg_per_transaction"`
	MonthOverMonth   float64 `json:"month_over_month_change"` // percentage
}

type SpendingPatterns struct {
	MostActiveDay   string             `json:"most_active_day"` // Monday, Tuesday, etc.
	PeakHour        int                `json:"peak_hour"`       // 0-23
	AvgWeekdaySpend float64            `json:"avg_weekday_spend"`
	AvgWeekendSpend float64            `json:"avg_weekend_spend"`
	ByDayOfWeek     map[string]float64 `json:"by_day_of_week"`
	ByHour          map[int]float64    `json:"by_hour"`
}

type TransactionAnomaly struct {
	TransactionID  uuid.UUID `json:"transaction_id"`
	Date           time.Time `json:"date"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Category       string    `json:"category"`
	ExpectedAmount float64   `json:"expected_amount"`
	Deviation      float64   `json:"deviation"` // multiplier (e.g., 3x normal)
}
