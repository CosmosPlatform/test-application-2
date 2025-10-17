package api

// Product represents a product in the system
type Product struct {
	ID          string  `json:"id" example:"product-123" description:"Unique product identifier"`
	Name        string  `json:"name" example:"Premium Widget" description:"Product name"`
	Description string  `json:"description" example:"A high-quality widget for all your needs" description:"Product description"`
	Price       float64 `json:"price" example:"29.99" description:"Product price"`
	Stock       int     `json:"stock" example:"100" description:"Available stock quantity"`
	Category    string  `json:"category" example:"Electronics" description:"Product category"`
}
