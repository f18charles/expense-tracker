package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/database"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepo struct {}

func NewBudgetRepo() *BudgetRepo {
	return &BudgetRepo{}
}

func (br *BudgetRepo) CreateBudget(budget *models.Budget) error {
	return database.DB.Create(budget).Error
}

func (br *BudgetRepo) GetBudgetByID(budget_id uuid.UUID) (*models.Budget, error) {
	var budget models.Budget
	result := database.DB.Where("id = ?", budget_id).First(&budget)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &budget, nil
}

func (br *BudgetRepo) UpdateBudget(budget *models.Budget) error {
	return database.DB.Save(budget).Error
}

func (br *BudgetRepo) ListBudgetsByUser(user_id uuid.UUID) ([]models.Budget, error) {
	budgets := []models.Budget{}
	result := database.DB.Where("user_id = ?", user_id).Find(&budgets)
	if result.Error != nil {
		return nil, result.Error
	}
	return budgets, nil
}

func (br *BudgetRepo) DeleteBudget(budget_id uuid.UUID) error {
	result := database.DB.Delete(&models.Budget{}, "id = ?", budget_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
