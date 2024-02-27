package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		var myName string = "Li Wei Chen"
		log.Printf("Request - Method: %s | Status: %d | Duration: %v | %s", c.Request.Method, c.Writer.Status(), duration, myName)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(LoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	router.GET("/bye", func(c *gin.Context) {
		c.String(200, "Good bye, to World")
	})

	authGroup := router.Group("/api")
	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Authenticated and authorized"})
		})

		authGroup.GET("/view_employee", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "username : forestDev"})
		})
	}

	router.Run(":8080")
}
