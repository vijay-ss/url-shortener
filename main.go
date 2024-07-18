package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test URL shortener",
		})
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("failed to start web server on port :9809, %v", err))
	}
}