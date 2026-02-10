package main

import (
	"net/http"
	// "github.com/f18charles/expense-tracker/internal/database"
	"fmt"
	"github.com/f18charles/expense-tracker/internal/api/handlers"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func main() {
	// static and assets


	// db check
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

	// normal routes
	http.HandleFunc("/", handlers.Dashboard)

	port := 5000
	fmt.Printf("Server is running on %v\n",port)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
