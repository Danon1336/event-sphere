package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Пропускаем OPTIONS-запросы (для CORS preflight)
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// Получаем или генерируем CSRF-токен
		csrfToken, err := c.Cookie("csrf_token")
		if err != nil || csrfToken == "" {
			tokenBytes := make([]byte, 32)
			_, err := rand.Read(tokenBytes)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSRF token"})
				c.Abort()
				return
			}
			csrfToken = base64.StdEncoding.EncodeToString(tokenBytes)
			c.SetCookie("csrf_token", csrfToken, 3600, "/", "", false, false)
		}

		// Для POST-запросов проверяем токен
		if c.Request.Method == "POST" {
			headerToken := c.GetHeader("X-CSRF-Token")
			if headerToken == "" || headerToken != csrfToken {
				c.JSON(http.StatusForbidden, gin.H{"error": "Invalid or missing CSRF token"})
				c.Abort()
				return
			}
		}

		// Устанавливаем токен в заголовок ответа
		c.Header("X-CSRF-Token", csrfToken)
		c.Next()
	}
}
