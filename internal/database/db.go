package database

import (
	"github.com/f18charles/expense-tracker/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// create database
func Setup(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Transaction{})
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("/internal/database/storage/app.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Transaction{})
	return db
}

// auth
func AddUser(name, email, password string) (string, error) {

	return "", nil
}
