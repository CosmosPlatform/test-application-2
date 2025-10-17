package api

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Element not found" description:"Error message"`
	Code    string `json:"code" example:"NOT_FOUND" description:"Error code"`
	Message string `json:"message" example:"The requested element could not be found" description:"Detailed error message"`
}
