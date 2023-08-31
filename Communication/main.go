package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/send-notification", func(c *gin.Context) {
		// Logic to send health-related notifications to users
		c.JSON(200, gin.H{"message": "Notification sent"})
	})

	router.Run(":8004")
}
