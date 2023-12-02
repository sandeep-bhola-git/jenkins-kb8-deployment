package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", HealthCheck())

	router.Run("localhost:7701")
}

func HealthCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		timeout := c.Param("toSecs")
		if timeout != "" {
			to, _ := time.ParseDuration(timeout + "s")
			time.Sleep(to)
		}
		c.JSON(201, gin.H{
			"status":        1,
			"response_data": "true",
		})
	}
}
