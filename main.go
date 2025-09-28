package main

import (
	"application/handlers"

	"github.com/gin-gonic/gin"
)


// @title Base application API
func main() {
	router := gin.Default()
	baseGroup := router.Group("/")

	handlers.RegisterHealthCheckRoutes(baseGroup)

	router.Run(":8080")
}
