package main

import (
	"net/http"
	// "github.com/f18charles/expense-tracker/internal/database"
	"fmt"
	"log"

	"github.com/f18charles/expense-tracker/internal/database"
	"github.com/f18charles/expense-tracker/web/app"
)

func main() {
	

	db := database.Init()
	r := app.Routes(db)

	port := 5000
	fmt.Printf("Server is running on %v\n",port)
	log.Fatal(http.ListenAndServe(":5000", r))

}
