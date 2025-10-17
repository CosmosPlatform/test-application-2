package handlers

import (
	"application/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

// In-memory storage for demo purposes
var products = map[string]*api.Product{
	"product-123": {
		ID:          "product-123",
		Name:        "Premium Widget",
		Description: "A high-quality widget for all your needs",
		Price:       29.99,
		Stock:       100,
		Category:    "Electronics",
	},
	"product-456": {
		ID:          "product-456",
		Name:        "Standard Widget",
		Description: "An affordable widget for everyday use",
		Price:       19.99,
		Stock:       250,
		Category:    "Electronics",
	},
	"product-789": {
		ID:          "product-789",
		Name:        "Deluxe Widget",
		Description: "The ultimate widget with premium features",
		Price:       49.99,
		Stock:       50,
		Category:    "Premium",
	},
}

func RegisterProductRoutes(e *gin.RouterGroup) {
	productsGroup := e.Group("/products")

	productsGroup.GET("/:id", getProductHandler)
}

// @Summary Get Product
// @Description Fetch product details by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} api.Product
// @Failure 404 {object} api.ErrorResponse
// @Router /products/{id} [get]
func getProductHandler(c *gin.Context) {
	productID := c.Param("id")

	product, exists := products[productID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "Product not found",
			Code:    "PRODUCT_NOT_FOUND",
			Message: "The requested product could not be found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
