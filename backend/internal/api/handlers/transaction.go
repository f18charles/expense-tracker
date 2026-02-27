package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/f18charles/expense-tracker/internal/api/middleware"
	"github.com/f18charles/expense-tracker/internal/models"
	"gorm.io/gorm"
)

func AddTransaction(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Redirect(w,r,"/add",http.StatusSeeOther)
			return
		}

		userID, err := middleware.GetUserID(r)
		if err != nil {
			http.Redirect(w,r, "/signin", http.StatusSeeOther)
			return
		}

		description := r.FormValue("description")
		category := r.FormValue("category")
		amount := r.FormValue("amount")
		extype := r.FormValue("type")

		money,err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			return
		}

		is_income := extype == "income"

		transactions := models.Transaction{
			Date: time.Now(),
			Description: description,
			Category: category,
			Amount: money,
			IsIncome: is_income,
			UserID: userID,
		}

		if err := db.Create(&transactions).Error; err != nil {
			http.Error(w,"Couldn't add transaction",http.StatusInternalServerError)
			return
		}
		// 7. Redirect back to dashboard
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}
