package services

import (
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
)

type AccCreateRequest struct {
	Name    string  `json:"name" binding:"required"`
	Type    string  `json:"type" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}

type AccountService struct {
	accountRepo *repository.AccountRepo
}

func NewAccService() *AccountService {
	return &AccountService{
		accountRepo: repository.NewAccountRepo(),
	}
}

type AccUpdateRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (as *AccountService) AccountCreate(user_id uuid.UUID, req AccCreateRequest) (*models.Account, error) {
	// create account
	account := &models.Account{
		Name:    req.Name,
		UserID:  user_id,
		Type:    req.Type,
		Balance: req.Balance,
	}

	// save
	if err := as.accountRepo.CreateAccount(account); err != nil {
		return nil, err
	}

	return account, nil
}

func (as *AccountService) GetAccount(userID, accID uuid.UUID) (*models.Account, error) {
	account, err := as.accountRepo.GetAccountByID(accID)
	if err != nil {
		return nil, err
	}

	if account.UserID != userID {
		return nil, utils.ErrForbidden
	}
	return account, nil
}

func (as *AccountService) AccountUpdate(userID, accID uuid.UUID, req AccUpdateRequest) (*models.Account, error) {
	account, err := as.accountRepo.GetAccountByID(accID)
	if err != nil {
		return nil, err
	}
	if account.UserID != userID {
		return nil, utils.ErrForbidden
	}
	if req.Name != "" {
		account.Name = req.Name
	}
	if req.Type != "" {
		account.Type = req.Type
	}
	if err := as.accountRepo.UpdateAccount(account); err != nil {
		return nil, err
	}
	return account, nil
}

func (as *AccountService) AccountList(user_id uuid.UUID) ([]models.Account, error) {
	accounts, err := as.accountRepo.ListAccountByUser(user_id)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (as *AccountService) AccountDelete(userID, accID uuid.UUID) error {
	account, err := as.accountRepo.GetAccountByID(accID)
	if err != nil {
		return err
	}

	if account.UserID != userID {
		return utils.ErrForbidden
	}

	return as.accountRepo.DeleteAccount(accID)

}
