package services

import (
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
)

type TxCreateRequest struct {
	CategoryID    *uuid.UUID `json:"category_id" binding:"required"`
	Amount        float64    `json:"amount" binding:"required"`
	Type          string     `json:"type" binding:"required"`
	Description   string     `json:"description" binding:"required"`
	PaymentMethod string     `json:"payment_method" binding:"required"`
	ReferenceID   string     `json:"reference_id" binding:"required"`
	Status        string     `json:"status" binding:"required"`
}

type TxService struct {
	txRepo *repository.TransactionRepo
}

func NewTxService() *TxService {
	return &TxService{
		txRepo: repository.NewTransactionRepo(),
	}
}

type TxUpdateRequest struct {
	Description string `json:"description"`
}

func (ts *TxService) TxCreate(user_id uuid.UUID, req TxCreateRequest) (*models.Transaction, error) {
	tx := &models.Transaction{
		CategoryID:    req.CategoryID,
		Amount:        req.Amount,
		Description:   req.Description,
		PaymentMethod: req.PaymentMethod,
		ReferenceID:   req.ReferenceID,
		Status:        req.Status,
	}

	if err := ts.txRepo.CreateTransaction(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

func (ts *TxService) TxGet(user_id, txID uuid.UUID) (*models.Transaction, error) {
	tx, err := ts.txRepo.GetTransactionByID(txID)
	if err != nil {
		return nil, err
	}
	if tx.UserID != user_id {
		return nil, utils.ErrForbidden
	}
	return tx, nil
}

func (ts *TxService) TxUpdate(user_id, txID uuid.UUID, req TxUpdateRequest) (*models.Transaction, error) {
	tx, err := ts.txRepo.GetTransactionByID(txID)
	if err != nil {
		return nil, err
	}
	if tx.UserID != user_id {
		return nil, utils.ErrForbidden
	}
	if req.Description != "" {
		tx.Description = req.Description
	}
	if err := ts.txRepo.UpdateTransaction(tx); err != nil {
		return nil, err
	}
	return tx, nil
}

func (ts *TxService) TxList(user_id uuid.UUID) ([]models.Transaction, error) {
	txs, err := ts.txRepo.ListTransactionsByUser(user_id)
	if err != nil {
		return nil, err
	}
	return txs, nil
}
