package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		// Handle user registration logic
		c.JSON(200, gin.H{"message": "User registered"})
	})

	router.POST("/login", func(c *gin.Context) {
		// Handle user login logic
		c.JSON(200, gin.H{"message": "User logged in"})
	})

	router.Run(":8001")
}
