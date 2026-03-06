package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ConfirmAuthedUser(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, errors.New("unauthorized")
	}
	id, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("invalid user")
	}
	return id, nil
}