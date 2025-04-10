package middleware

import (
	"crypto/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CSRFMiddleware generates and verifies CSRF tokens
func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate CSRF token if it's a GET request
		if c.Request.Method == "GET" {
			tokenBytes := make([]byte, 16)
			_, err := rand.Read(tokenBytes)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSRF token"})
				c.Abort()
				return
			}
			csrfToken := string(tokenBytes)
			c.SetCookie("csrf_token", csrfToken, 3600, "/", "", false, false)
			c.Header("X-CSRF-Token", csrfToken)
		}

		// Verify CSRF token for POST requests
		if c.Request.Method == "POST" {
			csrfFromCookie, err := c.Cookie("csrf_token")
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "CSRF token missing"})
				c.Abort()
				return
			}
			csrfFromHeader := c.GetHeader("X-CSRF-Token")
			if csrfFromHeader == "" || csrfFromHeader != csrfFromCookie {
				c.JSON(http.StatusForbidden, gin.H{"error": "Invalid CSRF token"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
