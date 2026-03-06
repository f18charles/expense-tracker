package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ConfirmAuthedUser extracts the authenticated user ID from the Gin context.
// Returns uuid.Nil and an error when the user is not authenticated or the
// value in context cannot be parsed as a UUID.
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
