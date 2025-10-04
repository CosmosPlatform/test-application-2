package handlers

import (
	"application/api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory storage for demo purposes
var orders = map[string]*api.Order{
	"order-123": {
		ID:          "order-123",
		CustomerID:  "customer-456",
		ProductName: "Premium Widget",
		Quantity:    2,
		Price:       29.99,
		Total:       59.98,
		Status:      "pending",
		CreatedAt:   time.Now().Add(-time.Hour),
		UpdatedAt:   time.Now().Add(-time.Hour),
	},
	"order-456": {
		ID:          "order-456",
		CustomerID:  "customer-789",
		ProductName: "Standard Widget",
		Quantity:    1,
		Price:       19.99,
		Total:       19.99,
		Status:      "paid",
		CreatedAt:   time.Now().Add(-2 * time.Hour),
		UpdatedAt:   time.Now().Add(-time.Minute * 30),
	},
}

func RegisterOrderRoutes(e *gin.RouterGroup) {
	ordersGroup := e.Group("/orders")

	ordersGroup.GET("/:id", getOrderHandler)
	ordersGroup.PUT("/:id", updateOrderHandler)
}

// @Summary Get Order
// @Description Fetch order details before charging
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} api.Order
// @Failure 404 {object} api.ErrorResponse
// @Router /orders/{id} [get]
func getOrderHandler(c *gin.Context) {
	orderID := c.Param("id")

	order, exists := orders[orderID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "Order not found",
			Code:    "ORDER_NOT_FOUND",
			Message: "The requested order could not be found",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary Update Order
// @Description Mark order as paid after successful charge
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body api.OrderUpdateRequest true "Order update data"
// @Success 200 {object} api.Order
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Router /orders/{id} [put]
/*
func updateOrderHandler(c *gin.Context) {
	orderID := c.Param("id")

	order, exists := orders[orderID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "Order not found",
			Code:    "ORDER_NOT_FOUND",
			Message: "The requested order could not be found",
		})
		return
	}

	var updateReq api.OrderUpdateRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "The request body is malformed or missing required fields",
		})
		return
	}

	// Update order status
	order.Status = updateReq.Status
	order.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, order)
}
*/