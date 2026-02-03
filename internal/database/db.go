package database

import (
	"encoding/json"
	"time"
	"github.com/f18charles/expense-tracker/internal/models"
)

const UserFile = "users.json"

func SaveUser(u models.User) error {
	users,_ LoadAllUsers()
	users = append(users, u)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(UserFile, data, 0644)
}

func LoadAllUsers() ([]models.User, error) {
	data, err := os.ReadFile(UserFile)
	if err != nil {
		return []models.User{}, nil
	}
	var users []models.User
	err = json.Unmarshal(data, &users)
	return users,err
}