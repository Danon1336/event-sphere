package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "List content"})
}

func GetSupport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Support content"})
}

func GetNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Notifications content"})
}

func GetFriends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Friends content"})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")
	c.JSON(http.StatusOK, gin.H{"id": userID, "role": role})
}

func AddEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Event added"})
}
