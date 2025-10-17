package handlers

import (
	"application/api"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory storage for demo purposes
var users = map[string]*api.User{
	"user-123": {
		ID:          "user-123",
		Email:       "john.doe@example.com",
		Username:    "johndoe",
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1234567890",
		Status:      "active",
		CreatedAt:   time.Now().Add(-30 * 24 * time.Hour), // 30 days ago
		UpdatedAt:   time.Now().Add(-2 * time.Hour),       // 2 hours ago
		Verified:    true,
	},
	"user-456": {
		ID:          "user-456",
		Email:       "jane.smith@example.com",
		Username:    "janesmith",
		FirstName:   "Jane",
		LastName:    "Smith",
		PhoneNumber: "+9876543210",
		Status:      "active",
		CreatedAt:   time.Now().Add(-60 * 24 * time.Hour), // 60 days ago
		UpdatedAt:   time.Now().Add(-24 * time.Hour),      // 1 day ago
		Verified:    true,
	},
	"user-789": {
		ID:        "user-789",
		Email:     "bob.wilson@example.com",
		Username:  "bobwilson",
		FirstName: "Bob",
		LastName:  "Wilson",
		Status:    "inactive",
		CreatedAt: time.Now().Add(-90 * 24 * time.Hour), // 90 days ago
		UpdatedAt: time.Now().Add(-7 * 24 * time.Hour),  // 7 days ago
		Verified:  false,
	},
}

func RegisterUserRoutes(e *gin.RouterGroup) {
	usersGroup := e.Group("/users")

	usersGroup.GET("/:id", getUserHandler)
}

// @Summary Get User
// @Description Fetch user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} api.User
// @Failure 404 {object} api.ErrorResponse
// @Router /users/{id} [get]
func getUserHandler(c *gin.Context) {
	userID := c.Param("id")

	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, api.ErrorResponse{
			Error:   "User not found",
			Code:    "USER_NOT_FOUND",
			Message: fmt.Sprintf("No user found with ID %s", userID),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
