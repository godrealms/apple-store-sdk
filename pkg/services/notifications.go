package services

import (
	"encoding/json"
	"net/http"
)

// NotificationHandler defines a function type for handling notifications
type NotificationHandler func(payload NotificationPayload) error

// NotificationHandlers maps notification types to handler functions
var NotificationHandlers = map[NotificationType]NotificationHandler{}

// RegisterNotificationHandler registers a handler for a specific notification type
func RegisterNotificationHandler(notificationType NotificationType, handler NotificationHandler) {
	NotificationHandlers[notificationType] = handler
}

// HandleNotification parses and processes the notification payload
func HandleNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload NotificationPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	handler, exists := NotificationHandlers[payload.NotificationType]
	if !exists {
		http.Error(w, "No handler for notification type", http.StatusNotImplemented)
		return
	}

	if err := handler(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification handled successfully"))
}

// Example default handlers
func handleRenewal(payload NotificationPayload) error {
	// Business logic for renewal notifications
	// e.g., update subscription status in the database
	return nil
}

func handleCancel(payload NotificationPayload) error {
	// Business logic for cancellation notifications
	// e.g., mark subscription as canceled in the database
	return nil
}

func init() {
	// Register default handlers
	RegisterNotificationHandler(NotificationTypeRenewal, handleRenewal)
	RegisterNotificationHandler(NotificationTypeCancel, handleCancel)
}
