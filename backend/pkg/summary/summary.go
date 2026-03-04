package summary

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"gorm.io/gorm"
)

type MonthlySummary struct {
	Month      time.Month
	Year       int
	OpeningBal float64
	TotalExp   float64
	CurrBal    float64
	ByCategory map[*models.Category]float64
}

func GetMonthlySummary(db *gorm.DB, userID uint, month time.Month, year int) (MonthlySummary, error) {
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	var txs []models.Transaction

	err := db.Where("user_id=? AND date >= ? and date < ?", userID, start, end).Find(&txs).Error
	if err != nil {
		return MonthlySummary{}, err
	}

	summary := MonthlySummary{
		Month:      month,
		Year:       year,
		ByCategory: make(map[*models.Category]float64),
	}

	var totalIncome float64
	var totalExpense float64

	for _, tx := range txs {
		if tx.Type == "income" {
			totalIncome += tx.Amount
		} else {
			totalExpense += tx.Amount
			summary.ByCategory[tx.Category] += tx.Amount
		}
	}

	var previousTxs []models.Transaction

	err = db.Where("user_id = ? AND date < ?", userID, start).Find(&previousTxs).Error
	if err != nil {
		return MonthlySummary{}, err
	}

	var opening float64

	for _, tx := range previousTxs {
		if tx.Type == "income" {
			opening += tx.Amount
		} else {
			opening -= tx.Amount
		}
	}

	summary.OpeningBal = opening
	summary.TotalExp = totalExpense
	summary.CurrBal = (opening + totalIncome) - totalExpense

	return summary, nil
}
