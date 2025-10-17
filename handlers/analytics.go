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

// In-memory storage for demo purposes - stores tracked events
var trackedEvents []api.AnalyticsEvent

func RegisterAnalyticsRoutes(e *gin.RouterGroup) {
	analyticsGroup := e.Group("/events")

	analyticsGroup.POST("", postEventHandler)
}

// @Summary Track Analytics Event
// @Description Track an analytics event (e.g., page views, clicks, conversions)
// @Tags Analytics
// @Accept json
// @Produce json
// @Param event body api.AnalyticsEvent true "Analytics event data"
// @Success 200 {object} api.AnalyticsEventResponse
// @Failure 400 {object} api.ErrorResponse
// @Router /events [post]
func postEventHandler(c *gin.Context) {
	var event api.AnalyticsEvent

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Error:   "Invalid request body",
			Code:    "INVALID_REQUEST",
			Message: "The request body is malformed or missing required fields",
		})
		return
	}

	// Validate required fields
	if event.EventType == "" {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Error:   "Missing event_type",
			Code:    "MISSING_FIELD",
			Message: "The event_type field is required",
		})
		return
	}

	// Set timestamp if not provided
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	// Store the event (in a real system, this would go to a database or analytics platform)
	trackedEvents = append(trackedEvents, event)

	// Generate a unique event ID
	eventID := generateEventID()

	response := api.AnalyticsEventResponse{
		Success:   true,
		EventID:   eventID,
		Timestamp: time.Now(),
		Message:   "Event tracked successfully",
	}

	c.JSON(http.StatusOK, response)
}

// generateEventID creates a unique event ID
func generateEventID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("event-%s", hex.EncodeToString(b))
}
