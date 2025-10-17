package main

import (
	"application/handlers"

	_ "application/docs"

	"github.com/gin-gonic/gin"
)

// @title Payment Service API
func main() {
	router := gin.Default()
	baseGroup := router.Group("/")

	handlers.RegisterHealthCheckRoutes(baseGroup)
	handlers.RegisterSwaggerRoutes(baseGroup)

	router.Run(":8080")
}
