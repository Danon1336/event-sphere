package handlers

import (
	"crypto/rand"
	"database/sql"
	"net/http"
	"project/backend/db"
	"project/backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your-secret-key") // Замени на безопасный ключ

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate salt"})
		return
	}
	salt := string(saltBytes)

	hashedPassword := db.HashPassword(user.Password, salt)

	_, err = db.DB.Exec("INSERT INTO users (name, surname, email, password, role, salt) VALUES (?, ?, ?, ?, ?, ?)",
		user.Name, user.Surname, user.Email, hashedPassword, user.Role, salt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	err := db.DB.QueryRow("SELECT id, name, surname, email, password, role, salt FROM users WHERE email = ?", input.Email).
		Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.Role, &user.Salt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	hashedInput := db.HashPassword(input.Password, user.Salt)
	if hashedInput != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", tokenString, 3600*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "role": user.Role})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func GetCSRFToken(c *gin.Context) {
	tokenBytes := make([]byte, 16)
	rand.Read(tokenBytes)
	token := string(tokenBytes)
	c.Header("X-CSRF-Token", token)
	c.JSON(http.StatusOK, gin.H{"csrf_token": token})
}
