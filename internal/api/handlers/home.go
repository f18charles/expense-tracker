package handlers

import (
	"net/http"
	"time"

	"github.com/f18charles/expense-tracker/internal/api/middleware"
	"github.com/f18charles/expense-tracker/internal/models"
	"github.com/f18charles/expense-tracker/internal/utils"
	"gorm.io/gorm"
)

func Home(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middleware.GetUserID(r)
		if err != nil {
			http.Redirect(w,r, "/signin", http.StatusSeeOther)
			return
		}
		
		now := time.Now()
		month := now.Month()
		year := now.Year()

		var user models.User

		data := struct {
			UserID uint
			UserName string
			Month time.Month
			Year int
		}{
			UserID: userID,
			UserName: user.Name,
			Month: month,
			Year: year,
		}

		if r.URL.Path != "/" {
			utils.RenderTemplate(w, "error_page.html", map[string]any{
				"Code":    http.StatusNotFound,
				"Message": "Route not found",
			})
			return
		}

		switch r.Method {
		case http.MethodGet:
			utils.RenderTemplate(w, "home.html", data)
		default:
			utils.RenderTemplate(w, "error_page.html", map[string]any{
				"Code":    http.StatusMethodNotAllowed,
				"Message": "Method not allowed",
			})
		}
	}
}