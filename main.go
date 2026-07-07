package main

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/roll", getRoll)

	router.Run("localhost:8080")
}

func getRoll(c *gin.Context) {
	sidesStr := c.Query("sides")
	sides, err := strconv.Atoi(sidesStr)

	if err != nil {
		c.JSON(400, gin.H{"message": "Sides must be a number!"})
		return
	}

	result := rand.Intn(sides) + 1
	c.IndentedJSON(200, gin.H{"result": result})
}
