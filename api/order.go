package api

import "time"

// Order represents an order in the system
type Order struct {
	ID          string    `json:"id" example:"order-123" description:"Unique order identifier"`
	CustomerID  string    `json:"customer_id" example:"customer-456" description:"Customer identifier"`
	ProductName string    `json:"product_name" example:"Premium Widget" description:"Name of the ordered product"`
	Quantity    int       `json:"quantity" example:"2" description:"Quantity of items ordered"`
	Price       float64   `json:"price" example:"29.99" description:"Price per item"`
	Total       float64   `json:"total" example:"59.98" description:"Total order amount"`
	Status      string    `json:"status" example:"pending" description:"Order status (pending, paid, shipped, delivered)"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z" description:"Order creation timestamp"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z" description:"Last update timestamp"`
}

// OrderUpdateRequest represents the request to update an order
type OrderUpdateRequest struct {
	Status string `json:"status" example:"paid" description:"New order status"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Order not found" description:"Error message"`
	Code    string `json:"code" example:"ORDER_NOT_FOUND" description:"Error code"`
	Message string `json:"message" example:"The requested order could not be found" description:"Detailed error message"`
}
