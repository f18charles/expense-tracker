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

func (ar *AccountRepo) GetAccountByID(accID uuid.UUID) (*models.Account, error) {
	var account models.Account
	result := database.DB.Where("id = ?", accID).First(&account)
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

func (ar *AccountRepo) ListAccountByUser(user_id uuid.UUID) ([]models.Account, error) {
	accounts := []models.Account{}
	result := database.DB.Where("user_id = ?", user_id).Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func (ar *AccountRepo) DeleteAccount(id uuid.UUID) error {
	result := database.DB.Delete(&models.Account{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
