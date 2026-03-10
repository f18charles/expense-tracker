package overview

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/google/uuid"
)

type DashboardOverview struct {
	NetWorth      float64              `json:"net_worth"`
	MonthlyBurn   float64              `json:"monthly_burn"`
	Accounts      []AccountBrief       `json:"accounts"`
	BudgetHealth  []BugdetBrief        `json:"budget_health"`
	GoalsProgress []GoalBrief          `json:"goals_progress"`
	RecentTx      []models.Transaction `json:"recent_transactions"`
	QuickInsights []string             `json:"quick_insights"`
}

type AccountBrief struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Balance  float64   `json:"balance"`
	Currency string    `json:"currency"`
}

type BugdetBrief struct {
	CategoryName string  `json:"category_name"`
	Spent        float64 `json:"spent"`
	Budget       float64 `json:"budget"`
	Percentage   float64 `json:"percentage"`
	Color        string  `json:"color"`
}

type GoalBrief struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	TargetAmount  float64    `json:"target_amount"`
	CurrentAmount float64    `json:"current_amount"`
	Percentage    float64    `json:"percentage"`
	Deadline      *time.Time `json:"deadline,omitempty"`
}
