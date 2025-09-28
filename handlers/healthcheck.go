package handlers

import (
	"application/models"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckRoutes(e *gin.RouterGroup) {
	healthCheckGroup := e.Group("/health")

	healthCheckGroup.GET("", healthCheckHandler)
}

// @Summary Health Check
// @Description Returns the health status of the application
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /health [get]
func healthCheckHandler(c *gin.Context) {
	response := models.HealthResponse{
		Status: "ok",
	}
	c.JSON(200, response)
}
