package models

// HealthResponse represents the response for health check endpoint
type HealthResponse struct {
	Status string `json:"status" example:"ok" description:"Health status of the application"`
}
