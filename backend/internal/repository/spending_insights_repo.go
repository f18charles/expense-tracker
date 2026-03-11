package repository

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpendingInsightsRepo struct {
	db *gorm.DB
}

func NewSpendingInsightsRepo(db *gorm.DB) *SpendingInsightsRepo {
	return &SpendingInsightsRepo{
		db: db,
	}
}

type QueryResult struct {
	models.Transaction
	CategoryName string
}

func (sir *SpendingInsightsRepo) GetPeriodExpenses(user_id uuid.UUID, start_date, end_date time.Time) ([]QueryResult, error) {
	var results []QueryResult
	if err := sir.db.Table("transactions").Select("transactions.*, categories.name as category_name").Joins("LEFT JOIN categories ON transactions.category_id = categories.id").Where("transactions.user_id = ? AND transactions.type = ? AND transactions.transaction_date BETWEEN ? AND ?", user_id, "expense", start_date, end_date).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
