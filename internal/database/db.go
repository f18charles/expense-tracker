package database

import (
	"gorm.io/gorm"
	"github.com/f18charles/expense-tracker/internal/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Transaction{})
}

