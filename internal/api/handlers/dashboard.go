package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/f18charles/expense-tracker/internal/api/middleware"
	utils "github.com/f18charles/expense-tracker/internal/utils"
	"github.com/f18charles/expense-tracker/pkg/summary"
	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middleware.GetUserID(r)
		if err != nil {
			http.Redirect(w,r, "/signin", http.StatusSeeOther)
			return
		}
		
		now := time.Now()
		month := now.Month()
		year := now.Year()

		sum, err := summary.GetMonthlySummary(db, userID, month, year)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := struct {
			UserID uint
			Summary summary.MonthlySummary
			Month time.Month
			Year int
		}{
			UserID: userID,
			Summary: sum,
			Month: month,
			Year: year,
		}

		if r.URL.Path != "/dashboard" {
			utils.RenderTemplate(w, "error_page.html", map[string]any{
				"Code":    http.StatusNotFound,
				"Message": "Route not found",
			})
			return
		}

		switch r.Method {
		case http.MethodGet:
			utils.RenderTemplate(w, "dash.html", data)
		default:
			utils.RenderTemplate(w, "error_page.html", map[string]any{
				"Code":    http.StatusMethodNotAllowed,
				"Message": "Method not allowed",
			})
		}
	}
}