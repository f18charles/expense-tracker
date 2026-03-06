package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger returns a Gin middleware that logs method, path, status and duration
// for each HTTP request — useful for development and debugging.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		log.Printf(
			"%s %s %v %s",
			c.Request.Method,
			path,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}
