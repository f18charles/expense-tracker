package handlers

import (
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/services"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	transactionService services.TxService
}

func NewTxHandler() *TransactionHandler {
	return &TransactionHandler{
		transactionService: *services.NewTxService(),
	}
}

func (th *TransactionHandler) ListTransactions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	id, ok := userID.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user")
		return
	}
	txs, err := th.transactionService.TxList(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "no transaction found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, txs)
}
func (th *TransactionHandler) CreateTransactions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	id, ok := userID.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user")
		return
	}
	var txreq services.TxCreateRequest
	if err := c.ShouldBindJSON(&txreq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tx, err := th.transactionService.TxCreate(id, txreq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create transaction")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, tx)
}

func (th *TransactionHandler) GetTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	id, ok := userID.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user")
		return
	}
	paramID := c.Param("id")
	txID, err := uuid.Parse(paramID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}
	tx, err := th.transactionService.TxGet(id, txID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "transactions not found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, tx)
}

func (th *TransactionHandler) UpdateTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	id := userID.(uuid.UUID)

	paramID := c.Param("id")
	txID, err := uuid.Parse(paramID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}

	var req services.TxUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := th.transactionService.TxUpdate(id, txID, req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, tx)
}

func (th *TransactionHandler) ExportTransactions(c *gin.Context) {}
