package app

import (
	"net/http"
	"github.com/f18charles/expense-tracker/internal/api/handlers"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	//auth
	mux.HandleFunc("/signup", handlers.SignUp(db))
	mux.HandleFunc("/signin", handlers.SignIn(db))
	mux.HandleFunc("/signout", handlers.SignOut())
	
	// static and assets
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	mux.HandleFunc("/", handlers.Home(db))
	mux.HandleFunc("/dashboard", handlers.Dashboard(db))
	mux.HandleFunc("/add", handlers.AddTransaction(db))

	return mux

}