package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"project/backend/db"
	"project/backend/handlers"
	"project/backend/middleware"
)

func main() {
	// Инициализация базы данных SQLite
	db.InitDB()

	// Настройка Gin
	r := gin.Default()

	// Настройка CORS для фронтенда
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"X-CSRF-Token"},
		AllowCredentials: true,
	}))

	// Применяем CSRF middleware ко всем маршрутам
	r.Use(middleware.CSRFMiddleware())

	// Публичные маршруты
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.POST("/logout", handlers.Logout)
	r.GET("/csrf-token", handlers.GetCSRFToken)

	// Защищенные маршруты (требуется JWT)
	protected := r.Group("/").Use(middleware.AuthMiddleware())
	{
		protected.GET("/list", handlers.GetList)
		protected.GET("/support", handlers.GetSupport)
		protected.GET("/notifications", handlers.GetNotifications)
		protected.GET("/friends", handlers.GetFriends)
		protected.GET("/profile", handlers.GetProfile)
		protected.POST("/events", middleware.RoleMiddleware("participant", "organizer", "moderator"), handlers.AddEvent)
	}

	// Запуск сервера
	r.Run(":8080")
}
