package main

import (
	"net/http"
	// "github.com/f18charles/expense-tracker/internal/database"
	"fmt"
	"log"

	"github.com/f18charles/expense-tracker/internal/api/handlers"
	"github.com/f18charles/expense-tracker/internal/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// static and assets

	db := database.Init()
	r := app.Routes(db)

	// db check
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

	database.Setup(db)

	// normal routes
	http.HandleFunc("/", handlers.Dashboard)

	port := 5000
	fmt.Printf("Server is running on %v\n",port)
	log.Fatal(http.ListenAndServe(":5000", r))

}
