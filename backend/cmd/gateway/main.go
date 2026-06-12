package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.JSON) {
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "nexus-gateway",
		})
	})

	log.Println("Gateway starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
