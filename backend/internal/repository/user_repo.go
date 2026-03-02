package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/database"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}
