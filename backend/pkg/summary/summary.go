package summary

import (
	"time"

	"github.com/f18charles/expense-tracker/internal/models"
	"gorm.io/gorm"
)

type MonthlySummary struct {
	Month      time.Month
	Year       int
	OpeningBal int64
	TotalExp   int64
	CurrBal    int64
	ByCategory map[string]int64
}

func GetMonthlySummary(db *gorm.DB,userID uint, month time.Month, year int) (MonthlySummary, error) {
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0,1,0)

	var txs []models.Transaction

	err := db.Where("user_id=? AND date >= ? and date < ?",userID,start,end).Find(&txs).Error
	if err != nil {
		return MonthlySummary{}, err
	}

	summary := MonthlySummary{
		Month: month, 
		Year: year,
		ByCategory: make(map[string]int64),
	}

	var totalIncome int64
	var totalExpense int64

	for _, tx := range txs {
		if tx.IsIncome {
			totalIncome += tx.Amount
		} else {
			totalExpense += tx.Amount
			summary.ByCategory[tx.Category]+=tx.Amount
		}
	}

	var previousTxs []models.Transaction

	err = db.Where("user_id = ? AND date < ?", userID, start).Find(&previousTxs).Error
	if err != nil {
		return MonthlySummary{}, err
	}
	
	var opening int64

	for _, tx := range previousTxs {
		if tx.IsIncome {
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
