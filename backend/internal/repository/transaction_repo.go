package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/database"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepo struct{}

func NewTransactionRepo() *TransactionRepo {
	return &TransactionRepo{}
}

func (tr *TransactionRepo) CreateTransaction(tx *models.Transaction) error {
	result := database.DB.Create(tx)
	return result.Error
}

func (tr *TransactionRepo) GetTransactionByID(txID uuid.UUID) (*models.Transaction, error) {
	var tx models.Transaction
	result := database.DB.Where("id = ?", txID).First(&tx)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &tx, nil
}

func (tr *TransactionRepo) UpdateTransaction(tx *models.Transaction) error {
	result := database.DB.Save(tx)
	return result.Error
}

func (tr *TransactionRepo) ListTransactionsByUser(userID uuid.UUID) ([]models.Transaction, error) {
	txs := []models.Transaction{}
	result := database.DB.Where("user_id = ?", userID).Find(&txs)
	if result.Error != nil {
		return nil, result.Error
	}
	return txs, nil
}

func (tr *TransactionRepo) DeleteTransaction(id uuid.UUID) error {
	result := database.DB.Delete(&models.Transaction{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
