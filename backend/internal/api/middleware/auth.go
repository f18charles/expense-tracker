package middleware

import (
	"net/http"
	"strings"

	"github.com/f18charles/piggy-bank/backend/internal/auth"
	"github.com/gin-gonic/gin"
)

// AuthRequired is a Gin middleware that enforces the presence of a valid
// `Authorization: Bearer <token>` header. On success it stores `user_id` in
// the request context for downstream handlers to use.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
