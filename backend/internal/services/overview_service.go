package services

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/pkg/overview"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OverviewService struct {
	db *gorm.DB
}

func NewOverviewService(db *gorm.DB) *OverviewService {
	return &OverviewService{
		db: db,
	}
}

func (os *OverviewService) GetDashboardOverview(user_id uuid.UUID) (*overview.DashboardOverview, error) {
	over_view := &overview.DashboardOverview{
		QuickInsights: []string{},
	}

	// get all accounts and networth
	var accounts []models.Account
	if err := os.db.Where("user_id = ?", user_id).Find(&accounts).Error; err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		over_view.NetWorth += acc.Balance
		over_view.Accounts = append(over_view.Accounts, overview.AccountBrief{
			ID:       acc.ID,
			Name:     acc.Name,
			Type:     acc.Type,
			Balance:  acc.Balance,
			Currency: acc.Currency,
		})
	}

	// calculate monthly burn
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	var monthly_expenses float64
	os.db.Model(&models.Transaction{}).Where("user_id = ? and type = ? and transaction_date = ?", user_id, "expense", startOfMonth).Select("COALESCE(SUM(amount), 0)").Scan(&monthly_expenses)
	over_view.MonthlyBurn = monthly_expenses

	// get budget health
	var budgets []models.Budget
	if err := os.db.Preload("Category").Where("user_id = ?", user_id).Find(&budgets).Error; err != nil {
		return nil, err
	}

	for _, b := range budgets {
		if b.CategoryID != uuid.Nil {
			percentage := (b.Spent / b.Amount) * 100
			over_view.BudgetHealth = append(over_view.BudgetHealth, overview.BugdetBrief{
				CategoryName: b.Category.Name,
				Spent:        b.Spent,
				Budget:       b.Amount,
				Percentage:   percentage,
				Color:        b.Category.Color,
			})

			// insights
			if percentage > 90 {
				over_view.QuickInsights = append(over_view.QuickInsights, "⚠️ You're close to your "+b.Category.Name+" budget limit")
			}
		}
	}

	// get goals progress
	var goals []models.Goal
	if err := os.db.Where("user_id = ?", user_id).Find(&goals).Error; err != nil {
		return nil, err
	}

	for _, g := range goals {
		percentage := (g.CurrentAmount / g.TargetAmount) * 100
		over_view.GoalsProgress = append(over_view.GoalsProgress, overview.GoalBrief{
			ID:            g.ID,
			Name:          g.Name,
			TargetAmount:  g.TargetAmount,
			CurrentAmount: g.CurrentAmount,
			Percentage:    percentage,
			Deadline:      g.Deadline,
		})
	}

	if err := os.db.Where("user_id = ?", user_id).Order("transaction_date DESC").Limit(5).Preload("Account").Preload("Category").Find(&over_view.RecentTx).Error; err != nil {
		return nil, err
	}

	// more insights
	if len(goals) > 0 {
		closestGoal := goals[0]
		for _, g := range goals {
			if g.Deadline != nil && (closestGoal.Deadline == nil || g.Deadline.Before(*closestGoal.Deadline)) {
				closestGoal = g
			}
		}
		if closestGoal.Deadline != nil {
			days_left := int(time.Until(*closestGoal.Deadline).Hours() / 24)
			over_view.QuickInsights = append(over_view.QuickInsights, "🎯 You have "+string(rune(days_left))+" days left for '"+closestGoal.Name+"'")
		}
	}
	return over_view, nil
}
