package storage

import (
	"encoding/csv"
	"expense-tracker/internal/models"
	"os"
	"strconv"
	"time"
)

func GetMonthlySummary(userID uint, month time.Month, year int) (models.MonthlySummary, error) {
	f, err := os.Open("transactions.csv")
	if err != nil {
		return models.MonthlySummary{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()

	summary := models.MonthlySummary{
		Month: month, Year: year,
		ByCategory: make(map[string]int64),
	}

	for _, r := range records {
		// r[1] is UserID, r[2] is Unix Date
		uID, _ := strconv.ParseUint(r[1], 10, 64)
		timestamp, _ := strconv.ParseInt(r[2], 10, 64)
		tTime := time.Unix(timestamp, 0)

		if uint(uID) == userID && tTime.Month() == month && tTime.Year() == year {
			amt, _ := strconv.ParseInt(r[5], 10, 64)
			isInc, _ := strconv.ParseBool(r[6])
			cat := r[4]

			if cat == "Opening Balance" {
				summary.OpeningBal = amt
			} else if isInc {
				summary.CurrBal += amt
			} else {
				summary.TotalExp += amt
				summary.ByCategory[cat] += amt
			}
		}
	}
	summary.CurrBal += (summary.OpeningBal - summary.TotalExp)
	return summary, nil
}
