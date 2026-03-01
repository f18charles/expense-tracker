package main

import (
	"log"

	"github.com/f18charles/piggy-bank/backend/internal/api"
	"github.com/f18charles/piggy-bank/backend/internal/config"
	"github.com/f18charles/piggy-bank/backend/internal/database"
)

func main() {
	// load config
	config.Load()

	// Connect to database
	database.Connect()

	r := api.SetupRouter()

	addr := ":" + config.App.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
