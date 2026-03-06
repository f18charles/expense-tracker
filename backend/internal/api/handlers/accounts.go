package handlers

import (
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/auth"
	"github.com/f18charles/piggy-bank/backend/internal/services"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccHandler() *AccountHandler {
	return &AccountHandler{
		accountService: *services.NewAccService(),
	}
}

// ListAccounts returns the list of accounts for the authenticated user.
func (ach *AccountHandler) ListAccounts(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	accounts, err := ach.accountService.AccountList(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "no accounts found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, accounts)
}

// CreateAccount creates a new account for the authenticated user.
func (ach *AccountHandler) CreateAccount(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var accreq services.AccCreateRequest
	if err := c.ShouldBindJSON(&accreq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	account, err := ach.accountService.AccountCreate(id, accreq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create account")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, account)
}

// GetAccount retrieves a single account by id for the authenticated user.
func (ach *AccountHandler) GetAccount(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	paramID := c.Param("id")
	accountID, err := uuid.Parse(paramID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}

	account, err := ach.accountService.GetAccount(id, accountID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "account not found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, account)
}

// UpdateAccount updates fields of an account owned by the authenticated user.
func (ach *AccountHandler) UpdateAccount(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	paramID := c.Param("id")
	accountID, err := uuid.Parse(paramID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}

	var req services.AccUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := ach.accountService.AccountUpdate(id, accountID, req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, account)
}

// DeleteAccount removes an account owned by the authenticated user.
func (ach *AccountHandler) DeleteAccount(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userID := id

	paramID := c.Param("id")
	accountID, err := uuid.Parse(paramID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
		return
	}

	err = ach.accountService.AccountDelete(userID, accountID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "deleted"})
}
