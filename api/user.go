package api

import "time"

// User represents a user in the system
type User struct {
	ID          string    `json:"id" example:"user-123" description:"Unique user identifier"`
	Email       string    `json:"email" example:"user@example.com" description:"User email address"`
	Username    string    `json:"username" example:"johndoe" description:"Username"`
	FirstName   string    `json:"first_name" example:"John" description:"User's first name"`
	LastName    string    `json:"last_name" example:"Doe" description:"User's last name"`
	PhoneNumber string    `json:"phone_number,omitempty" example:"+1234567890" description:"User's phone number"`
	Status      string    `json:"status" example:"active" description:"User status (active, inactive, suspended)"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z" description:"Account creation timestamp"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z" description:"Last update timestamp"`
	Verified    bool      `json:"verified" example:"true" description:"Whether the user's email is verified"`
}
