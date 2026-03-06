package services

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
)

type BudgetServices struct {
	budgetRepo *repository.BudgetRepo
}

func NewBudgetRepo() *BudgetServices {
	return &BudgetServices{
		budgetRepo: repository.NewBudgetRepo(),
	}
}

type BudgetCreateRequest struct {
	CategoryID *uuid.UUID `json:"category_id" binding:"required"`
	Amount float64 `json:"amount"`
	Spent float64 `json:"spent"`
	Period string `json:"period" binding:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time	`json:"end_date"`
} 

type BudgetUpdateRequest struct {
	CategoryID *uuid.UUID `json:"category_id"`
	Amount float64 `json:"amount"`
	Spent float64 `json:"spent"`
	Period string `json:"period"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time	`json:"end_date"`
}

func (bs *BudgetServices) BudgetCreate(user_id uuid.UUID, req BudgetCreateRequest) (*models.Budget, error) {
	budget := &models.Budget{
		UserID: user_id,
		CategoryID: *req.CategoryID,
		Amount: req.Amount,
		Spent: req.Amount,
		Period: req.Period,
		StartDate: *req.StartDate,
		EndDate: *req.EndDate,
	}
	if err := bs.budgetRepo.CreateBudget(budget); err != nil {
		return nil, err
	}
	return budget, nil
}

func (bs *BudgetServices) BudgetGet(user_id, budget_id uuid.UUID) (*models.Budget, error) {
	budget, err := bs.budgetRepo.GetBudgetByID(budget_id)
	if err != nil {
		return nil, err
	}
	if budget.UserID != user_id {
		return nil, utils.ErrForbidden
	}
	return budget, nil
}

func (bs *BudgetServices) BudgetUpdate(budget_id, user_id uuid.UUID, req BudgetUpdateRequest) (*models.Budget, error) {
	budget, err := bs.budgetRepo.GetBudgetByID(budget_id)
	if err != nil {
		return nil, err
	}
	if budget.UserID != user_id {
		return nil, utils.ErrForbidden
	}
	if req.CategoryID != nil {budget.CategoryID = *req.CategoryID}
	if req.Amount != 0 {budget.Amount = req.Amount}
	if req.Spent != 0 {budget.Spent = req.Spent}
	if req.Period != "" {budget.Period = req.Period}
	if req.StartDate != nil {budget.StartDate = *req.StartDate}
	if req.EndDate != nil {budget.EndDate = *req.EndDate}
	return budget, nil
}

func (bs *BudgetServices) BudgetList(user_id uuid.UUID) ([]models.Budget, error) {
	budgets, err := bs.budgetRepo.ListBudgetsByUser(user_id)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (bs *BudgetServices) BudgetDelete(budget_id uuid.UUID) error {
	if err := bs.budgetRepo.DeleteBudget(budget_id); err != nil {
		return err
	}
	return nil
}
