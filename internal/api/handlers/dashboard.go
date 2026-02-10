package handlers

import (
	"net/http"
	utils"github.com/f18charles/expense-tracker/internal/utils"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.RenderTemplate(w, "error_page.html", map[string]any{
			"Code":    http.StatusNotFound,
			"Message": "Route not found",
		})
		return
	}

	switch r.Method {
	case http.MethodGet:
		utils.RenderTemplate(w, "dash.html", nil)
	default:
		utils.RenderTemplate(w, "error_page.html", map[string]any{
			"Code":    http.StatusMethodNotAllowed,
			"Message": "Method not allowed",
		})
	}
}