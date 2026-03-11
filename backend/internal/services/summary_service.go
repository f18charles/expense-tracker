package services

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/f18charles/piggy-bank/backend/pkg/summary"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SummaryService struct {
	db           *gorm.DB
	summaryRepo  *repository.SummaryRepo
	tx_repo      *repository.TransactionRepo
	budget_repo  *repository.BudgetRepo
	account_repo *repository.AccountRepo
}

func NewSummaryService(db *gorm.DB) *SummaryService {
	return &SummaryService{
		db:           db,
		summaryRepo:  repository.NewSummaryRepo(db),
		tx_repo:      repository.NewTransactionRepo(),
		budget_repo:  repository.NewBudgetRepo(),
		account_repo: repository.NewAccountRepo(),
	}
}

func (s *SummaryService) GetMonthlySummary(user_id uuid.UUID, year int, month time.Month) (*summary.MonthlySummary, error) {
	// Get all transactions for the month
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	transactions, err := s.summaryRepo.GetTransactions(user_id, startDate, endDate)
	if err != nil {
		return nil, utils.ErrNotFound
	}

	// Get budgets for the month
	budgets, err := s.budget_repo.ListBudgetsByUser(user_id)
	if err != nil {
		return nil, err
	}
	budgetMap := make(map[uuid.UUID]float64)
	for _, b := range budgets {
		budgetMap[b.CategoryID] = b.Amount
	}

	categories, err := s.summaryRepo.GetCategories(user_id)
	if err != nil {
		return nil, utils.ErrNotFound
	}

	categoryMap := make(map[uuid.UUID]models.Category)
	for _, c := range categories {
		categoryMap[c.ID] = c
	}

	// Build summary
	mon_summary := &summary.MonthlySummary{
		UserID:     user_id,
		Year:       year,
		Month:      month,
		ByCategory: make(map[string]summary.CategorySpend),
	}

	categorySpends := make(map[uuid.UUID]float64)

	for _, tx := range transactions {
		if tx.Type == "income" {
			mon_summary.Income += tx.Amount
		} else {
			mon_summary.Expenses += tx.Amount
			if tx.CategoryID != nil {
				categorySpends[*tx.CategoryID] += tx.Amount
			}
		}
	}

	mon_summary.Savings = mon_summary.Income - mon_summary.Expenses
	if mon_summary.Income > 0 {
		mon_summary.SavingsRate = (mon_summary.Savings / mon_summary.Income) * 100
	}

	// Calculate category breakdowns
	totalExpenses := mon_summary.Expenses
	for catID, spent := range categorySpends {
		category := categoryMap[catID]
		percentage := 0.0
		if totalExpenses > 0 {
			percentage = (spent / totalExpenses) * 100
		}

		mon_summary.ByCategory[category.Name] = summary.CategorySpend{
			CategoryID:    catID,
			CategoryName:  category.Name,
			CategoryColor: category.Color,
			Spent:         spent,
			Budget:        budgetMap[catID],
			Percentage:    percentage,
		}
	}

	// Get previous month for comparison
	prevMonth := month - 1
	prevYear := year
	if prevMonth == 0 {
		prevMonth = 12
		prevYear--
	}

	prevmon_Summary, _ := s.GetMonthlySummary(user_id, prevYear, prevMonth)
	mon_summary.PreviousMonth = prevmon_Summary

	return mon_summary, nil
}

// GetYearlySummary aggregates monthly summaries for a year
func (s *SummaryService) GetYearlySummary(user_id uuid.UUID, year int) ([]summary.MonthlySummary, error) {
	var summaries []summary.MonthlySummary

	for month := time.January; month <= time.December; month++ {
		summary, err := s.GetMonthlySummary(user_id, year, month)
		if err != nil {
			continue // Skip months with no data
		}
		summaries = append(summaries, *summary)
	}

	return summaries, nil
}
