package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/database"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepo struct{}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{}
}

func (ar *AccountRepo) CreateAccount(account *models.Account) error {
	result := database.DB.Create(account)
	return result.Error
}

func (ar *AccountRepo) GetAccountByUser(user_id uuid.UUID) (*models.Account, error) {
	var account models.Account
	result := database.DB.Where("user_id = ?", user_id).First(&account)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &account, nil
}

func (ar *AccountRepo) GetAccountByID(id uuid.UUID) (*models.Account, error) {
	var account models.Account
	result := database.DB.Where("id = ?", id).First(&account)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &account, nil
}

func (ar *AccountRepo) UpdateAccount(account *models.Account) error {
	result := database.DB.Save(account)
	return result.Error
}