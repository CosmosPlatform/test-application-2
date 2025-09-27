package handlers

import "github.com/gin-gonic/gin"

func RegisterHealthCheckRoutes(e *gin.RouterGroup) {
	healthCheckGroup := e.Group("/health")

	healthCheckGroup.GET("", healthCheckHandler)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
