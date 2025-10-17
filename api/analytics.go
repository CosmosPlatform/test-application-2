package api

import "time"

// AnalyticsEvent represents an analytics event to be tracked
type AnalyticsEvent struct {
	EventType  string                 `json:"event_type" example:"page_view" description:"Type of the event"`
	UserID     string                 `json:"user_id" example:"user-123" description:"User identifier"`
	SessionID  string                 `json:"session_id" example:"session-456" description:"Session identifier"`
	Timestamp  time.Time              `json:"timestamp" example:"2023-01-01T00:00:00Z" description:"Event timestamp"`
	Properties map[string]interface{} `json:"properties,omitempty" description:"Additional event properties"`
}

// AnalyticsEventResponse represents the response after tracking an event
type AnalyticsEventResponse struct {
	Success   bool      `json:"success" example:"true" description:"Whether the event was successfully tracked"`
	EventID   string    `json:"event_id" example:"event-789" description:"Unique identifier for the tracked event"`
	Timestamp time.Time `json:"timestamp" example:"2023-01-01T00:00:00Z" description:"Server timestamp when event was received"`
	Message   string    `json:"message" example:"Event tracked successfully" description:"Response message"`
}
