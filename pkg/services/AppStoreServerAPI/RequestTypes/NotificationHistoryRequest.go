package RequestTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

type NotificationHistoryRequest struct {
	// Required. The start date of the timespan for the requested App Store Server Notification history records.
	// The startDate needs to precede the endDate. Choose a startDate that’s within the past 180 days from the current date.
	StartDate types.Timestamp `json:"startDate"`
	// Required. The end date of the timespan for the requested App Store Server Notification history records.
	// Choose an endDate that’s later than the startDate.
	// If you choose an endDate in the future, the endpoint automatically uses the current date as the endDate.
	EndDate types.Timestamp `json:"endDate"`
	// Optional. A notification type. Provide this field to limit the notification history records to those with this one notification type.
	// For a list of notifications types, see notificationType.
	// Note: You may include either the transactionId or the notificationType property (or neither) in your query, but not both.
	NotificationType types.NotificationType `json:"notificationType,omitempty"`
	// Optional. A notification subtype.
	// Provide this field to limit the notification history records to those with this one notification subtype.
	// For a list of subtypes, see subtype. If you specify a notificationSubtype,
	// you need to also specify its related notificationType.
	NotificationSubtype types.NotificationSubtype `json:"notificationSubtype,omitempty"`
	// Optional. A Boolean value you set to true to request only the notifications that haven’t reached your server successfully.
	// The response also includes notifications that the App Store server is currently retrying to send to your server.
	OnlyFailures types.OnlyFailures `json:"onlyFailures,omitempty"`
	// Optional. The transaction identifier, which may be an original transaction identifier,
	// of any transaction belonging to the customer.
	// Provide this field to limit the notification history request to this one customer.
	// Note: You may include either the transactionId or the notificationType property (or neither) in your query,
	// but not both.
	TransactionId types.TransactionId `json:"transactionId,omitempty"`
}
