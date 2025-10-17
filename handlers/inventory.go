package handlers

import (
	"application/api"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory storage for demo purposes
var inventory = map[string]*api.InventoryItem{
	"product-123": {
		ProductID:       "product-123",
		Available:       100,
		Reserved:        5,
		Total:           105,
		LastUpdated:     time.Now(),
		WarehouseID:     "warehouse-001",
		ReorderPoint:    20,
		ReorderQuantity: 50,
	},
	"product-456": {
		ProductID:       "product-456",
		Available:       250,
		Reserved:        10,
		Total:           260,
		LastUpdated:     time.Now(),
		WarehouseID:     "warehouse-001",
		ReorderPoint:    50,
		ReorderQuantity: 100,
	},
	"product-789": {
		ProductID:       "product-789",
		Available:       50,
		Reserved:        2,
		Total:           52,
		LastUpdated:     time.Now(),
		WarehouseID:     "warehouse-002",
		ReorderPoint:    10,
		ReorderQuantity: 30,
	},
}

var inventoryEvents []api.InventoryEvent

func RegisterInventoryRoutes(e *gin.RouterGroup) {
	e.POST("/events", postInventoryEventHandler)

	inventoryGroup := e.Group("/inventory")
	{
		inventoryGroup.POST("/reserve", reserveInventoryHandler)
		inventoryGroup.POST("/release", releaseInventoryHandler)
		inventoryGroup.GET("/:productId", getInventoryHandler)
	}
}

// @Summary Get Inventory
// @Description Fetch inventory details for a specific product
// @Tags Inventory
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} api.InventoryItem
// @Failure 404 {object} api.ErrorResponse
// @Router /inventory/{productId} [get]
func getInventoryHandler(c *gin.Context) {
	productID := c.Param("productId")

	item, exists := inventory[productID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "Inventory not found",
			Code:    "INVENTORY_NOT_FOUND",
			Message: fmt.Sprintf("No inventory found for product %s", productID),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Reserve Inventory
// @Description Reserve inventory for an order
// @Tags Inventory
// @Accept json
// @Produce json
// @Param request body api.ReserveInventoryRequest true "Reserve inventory request"
// @Success 200 {object} api.InventoryResponse
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 409 {object} api.ErrorResponse
// @Router /inventory/reserve [post]
func reserveInventoryHandler(c *gin.Context) {
	var request api.ReserveInventoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "The request body is malformed or missing required fields",
		})
		return
	}

	item, exists := inventory[request.ProductID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "Product not found",
			Code:    "PRODUCT_NOT_FOUND",
			Message: fmt.Sprintf("No inventory found for product %s", request.ProductID),
		})
		return
	}

	// Check if sufficient inventory is available
	if item.Available < request.Quantity {
		c.JSON(http.StatusConflict, api.ErrorResponse{
			Error:   "Insufficient inventory",
			Code:    "INSUFFICIENT_INVENTORY",
			Message: fmt.Sprintf("Only %d units available, but %d requested", item.Available, request.Quantity),
		})
		return
	}

	// Reserve the inventory
	item.Available -= request.Quantity
	item.Reserved += request.Quantity
	item.LastUpdated = time.Now()

	// Generate reservation ID if not provided
	reservationID := request.ReservationID
	if reservationID == "" {
		reservationID = generateReservationID()
	}

	response := api.InventoryResponse{
		Success:       true,
		ProductID:     request.ProductID,
		Available:     item.Available,
		Reserved:      item.Reserved,
		ReservationID: reservationID,
		Message:       "Inventory reserved successfully",
		Timestamp:     time.Now(),
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Release Inventory
// @Description Release reserved inventory back to available stock
// @Tags Inventory
// @Accept json
// @Produce json
// @Param request body api.ReleaseInventoryRequest true "Release inventory request"
// @Success 200 {object} api.InventoryResponse
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 409 {object} api.ErrorResponse
// @Router /inventory/release [post]
func releaseInventoryHandler(c *gin.Context) {
	var request api.ReleaseInventoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "The request body is malformed or missing required fields",
		})
		return
	}

	item, exists := inventory[request.ProductID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "Product not found",
			Code:    "PRODUCT_NOT_FOUND",
			Message: fmt.Sprintf("No inventory found for product %s", request.ProductID),
		})
		return
	}

	// Check if sufficient reserved inventory exists
	if item.Reserved < request.Quantity {
		c.JSON(http.StatusConflict, api.ErrorResponse{
			Error:   "Insufficient reserved inventory",
			Code:    "INSUFFICIENT_RESERVED",
			Message: fmt.Sprintf("Only %d units reserved, but %d requested for release", item.Reserved, request.Quantity),
		})
		return
	}

	// Release the inventory
	item.Reserved -= request.Quantity
	item.Available += request.Quantity
	item.LastUpdated = time.Now()

	response := api.InventoryResponse{
		Success:   true,
		ProductID: request.ProductID,
		Available: item.Available,
		Reserved:  item.Reserved,
		Message:   "Inventory released successfully",
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Track Inventory Event
// @Description Track an inventory-related event (e.g., stock updates, audits, movements)
// @Tags Inventory
// @Accept json
// @Produce json
// @Param event body api.InventoryEvent true "Inventory event data"
// @Success 200 {object} api.InventoryEventResponse
// @Failure 400 {object} api.ErrorResponse
// @Router /events [post]
func postInventoryEventHandler(c *gin.Context) {
	var event api.InventoryEvent

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "The request body is malformed or missing required fields",
		})
		return
	}

	// Set timestamp if not provided
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	// Store the event (in a real system, this would go to a database or event log)
	inventoryEvents = append(inventoryEvents, event)

	// Generate a unique event ID
	eventID := generateEventID()

	response := api.InventoryEventResponse{
		Success:   true,
		EventID:   eventID,
		Timestamp: time.Now(),
		Message:   "Inventory event tracked successfully",
	}

	c.JSON(http.StatusOK, response)
}

// Helper functions

func generateReservationID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("res-%s", hex.EncodeToString(b))
}

func generateEventID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("event-%s", hex.EncodeToString(b))
}
