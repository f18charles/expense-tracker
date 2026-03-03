package services

import (
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
)

type AccCreateRequest struct {
	Name      string    `json:"name" binding:"required"`
	Type      string    `json:"type" binding:"required"`
	Balance   float64   `json:"balance" binding:"required"`
	Currency  string    `json:"currency" binding:"required"`
}

type AccountService struct {
	accountRepo *repository.AccountRepo
}

func NewAccService() *AccountService {
	return &AccountService{
		accountRepo: repository.NewAccountRepo(),
	}
}

func (as *AccountService) CreateAccout() (*models.Account, error) {
	
}