package api

import "time"

// InventoryItem represents the inventory for a product
type InventoryItem struct {
	ProductID       string    `json:"product_id" example:"product-123" description:"Unique product identifier"`
	Available       int       `json:"available" example:"100" description:"Available stock quantity"`
	Reserved        int       `json:"reserved" example:"5" description:"Reserved stock quantity"`
	Total           int       `json:"total" example:"105" description:"Total stock quantity"`
	LastUpdated     time.Time `json:"last_updated" example:"2023-01-01T00:00:00Z" description:"Last inventory update timestamp"`
	WarehouseID     string    `json:"warehouse_id" example:"warehouse-001" description:"Warehouse identifier"`
	ReorderPoint    int       `json:"reorder_point" example:"20" description:"Reorder threshold"`
	ReorderQuantity int       `json:"reorder_quantity" example:"50" description:"Quantity to reorder when below threshold"`
}

// ReserveInventoryRequest represents a request to reserve inventory
type ReserveInventoryRequest struct {
	ProductID     string `json:"product_id" example:"product-123" description:"Product identifier" binding:"required"`
	Quantity      int    `json:"quantity" example:"2" description:"Quantity to reserve" binding:"required,min=1"`
	OrderID       string `json:"order_id" example:"order-456" description:"Order identifier"`
	ReservationID string `json:"reservation_id,omitempty" example:"res-789" description:"Optional reservation identifier"`
}

// ReleaseInventoryRequest represents a request to release reserved inventory
type ReleaseInventoryRequest struct {
	ProductID     string `json:"product_id" example:"product-123" description:"Product identifier" binding:"required"`
	Quantity      int    `json:"quantity" example:"2" description:"Quantity to release" binding:"required,min=1"`
	ReservationID string `json:"reservation_id" example:"res-789" description:"Reservation identifier"`
}

// InventoryResponse represents a response for inventory operations
type InventoryResponse struct {
	Success       bool      `json:"success" example:"true" description:"Whether the operation was successful"`
	ProductID     string    `json:"product_id" example:"product-123" description:"Product identifier"`
	Available     int       `json:"available" example:"98" description:"Available stock after operation"`
	Reserved      int       `json:"reserved" example:"7" description:"Reserved stock after operation"`
	ReservationID string    `json:"reservation_id,omitempty" example:"res-789" description:"Reservation identifier (for reserve operations)"`
	Message       string    `json:"message" example:"Inventory reserved successfully" description:"Response message"`
	Timestamp     time.Time `json:"timestamp" example:"2023-01-01T00:00:00Z" description:"Operation timestamp"`
}

// InventoryEvent represents an inventory event to be tracked
type InventoryEvent struct {
	EventType  string                 `json:"event_type" example:"stock_update" description:"Type of inventory event" binding:"required"`
	ProductID  string                 `json:"product_id" example:"product-123" description:"Product identifier"`
	Quantity   int                    `json:"quantity" example:"10" description:"Quantity involved in the event"`
	Timestamp  time.Time              `json:"timestamp" example:"2023-01-01T00:00:00Z" description:"Event timestamp"`
	Properties map[string]interface{} `json:"properties,omitempty" description:"Additional event properties"`
}

// InventoryEventResponse represents the response after tracking an inventory event
type InventoryEventResponse struct {
	Success   bool      `json:"success" example:"true" description:"Whether the event was successfully tracked"`
	EventID   string    `json:"event_id" example:"event-789" description:"Unique identifier for the tracked event"`
	Timestamp time.Time `json:"timestamp" example:"2023-01-01T00:00:00Z" description:"Server timestamp when event was received"`
	Message   string    `json:"message" example:"Event tracked successfully" description:"Response message"`
}
