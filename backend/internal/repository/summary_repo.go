package repository

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SummaryRepo struct {
	db *gorm.DB
}

func NewSummaryRepo(db *gorm.DB) *SummaryRepo {
	return &SummaryRepo{
		db: db,
	}
}

func (sr *SummaryRepo) GetTransactions(user_id uuid.UUID, start_date, end_date time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := sr.db.Where("user_id = ? AND transaction_date >= ? AND transaction_date < ?",
		user_id, start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (sr *SummaryRepo) GetCategories(user_id uuid.UUID) ([]models.Category, error) {
	var categories []models.Category
	if err := sr.db.Where("user_id = ? OR is_default = true", user_id).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
