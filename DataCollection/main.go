package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/submit-data", func(c *gin.Context) {
		// Handle health data submission logic
		c.JSON(200, gin.H{"message": "Data submitted"})
	})

	router.Run(":8002")
}
