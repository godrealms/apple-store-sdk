package _ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

type ResponseBodyV2DecodedPayload struct {
	// The in-app purchase event for which the App Store sends this version 2 notification.
	NotificationType types.NotificationType `json:"notificationType"`
	// Additional information that identifies the notification event.
	// The subtype field is present only for specific version 2 notifications.
	Subtype types.Subtype `json:"subtype"`
	// The object that contains the app metadata and signed renewal and transaction information.
	// The data, summary, and externalPurchaseToken fields are mutually exclusive.
	// The payload contains only one of these fields.
	Data Data `json:"data"`
	// The summary data that appears when the App Store server completes your request to extend a subscription renewal date for eligible subscribers.
	// For more information, see Extend Subscription Renewal Dates for All Active Subscribers.
	// The data, summary, and externalPurchaseToken fields are mutually exclusive.
	// The payload contains only one of these fields.
	Summary types.Summary `json:"summary"`
	// This field appears when the notificationType is ExternalPurchaseToken.
	// The data, summary, and externalPurchaseToken fields are mutually exclusive.
	// The payload contains only one of these fields.
	ExternalPurchaseToken externalPurchaseToken `json:"externalPurchaseToken"`
	// The App Store AppStoreServerAPI Notification version number, "2.0".
	Version types.Version `json:"version"`
	// The UNIX time, in milliseconds, that the App Store signed the JSON Web Signature data.
	SignedDate types.Timestamp `json:"signedDate"`
	// A unique identifier for the notification. Use this value to identify a duplicate notification.
	NotificationUUID types.UUID `json:"notificationUUID"`
}
