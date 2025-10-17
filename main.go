package main

import (
	"application/handlers"

	_ "application/docs"

	"github.com/gin-gonic/gin"
)

// @title Analytics Service API
func main() {
	router := gin.Default()
	baseGroup := router.Group("/")

	handlers.RegisterHealthCheckRoutes(baseGroup)
	handlers.RegisterAnalyticsRoutes(baseGroup)
	handlers.RegisterSwaggerRoutes(baseGroup)

	router.Run(":8080")
}
