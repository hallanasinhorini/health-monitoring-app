package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/recommendations/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		// Logic to generate health recommendations for the user with userID
		recommendations := []string{"Exercise more", "Eat more fruits"}
		c.JSON(200, recommendations)
	})

	router.Run(":8003")
}
