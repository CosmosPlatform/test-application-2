package handlers

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwaggerRoutes(e *gin.RouterGroup) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
