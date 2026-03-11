package repository

import (
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

