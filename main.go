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

	if sides <= 0 {
		c.JSON(400, gin.H{"message": "Sides must not be zero or lower"})
		return
	}

	result := rand.Intn(sides) + 1
	average := float64(1+sides) / 2
	probability := 1 / float64(sides)

	c.IndentedJSON(200, gin.H{
		"result":      result,
		"sides":       sides,
		"average":     average,
		"probability": probability,
	})
}
