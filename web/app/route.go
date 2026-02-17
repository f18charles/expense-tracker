package app

import (
	"net/http"
	"github.com/f18charles/expense-tracker/internal/api/handlers"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/dashboard", handlers.Dashboard(db))
	// mux.HandleFunc("/add", handlers.AddTransaction(db))

	return mux

}