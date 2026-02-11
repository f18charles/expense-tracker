package main

import (
	"fmt"
	"github.com/f18charles/expense-tracker/internal/api/handlers"
	"github.com/f18charles/expense-tracker/internal/database"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"net/http"
)

func main() {
	// static and assets

	//database
	db, err := gorm.Open(sqlite.Open("internal/database/database.db"))
	if err != nil {
		panic("Cannot connect to database")
	}
	database.Migrate(db)

	// normal routes
	http.HandleFunc("/", handlers.Dashboard)

	port := 5000
	fmt.Printf("Server is running on %v\n", port)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
