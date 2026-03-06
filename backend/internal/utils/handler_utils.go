package utils

import (
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ConfirmAuthedUser(c *gin.Context) (uuid.UUID) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return uuid.Nil
	}
	id, ok := userID.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user")
		return uuid.Nil
	}
	return id
}