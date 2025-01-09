package models

// NotificationHistoryRequest The request body for notification history.
type NotificationHistoryRequest struct {
	// Required. The start date of the timespan for the requested App Store Server Notification history records.
	// The startDate needs to precede the endDate. Choose a startDate that’s within the past 180 days from the current date.
	StartDate startDate `json:"startDate"`

	// Required. The end date of the timespan for the requested App Store Server Notification history records.
	// Choose an endDate that’s later than the startDate. If you choose an endDate in the future,
	// the endpoint automatically uses the current date as the endDate.
	EndDate endDate `json:"endDate"`

	// Optional. A notification type. Provide this field to limit the notification history records to those with this one notification type.
	// For a list of notifications types, see notificationType.
	// Note: You may include either the transactionId or the notificationType property (or neither) in your query, but not both.
	NotificationType notificationType `json:"notificationType"`

	// Optional. A notification subtype.
	// Provide this field to limit the notification history records to those with this one notification subtype.
	// For a list of subtypes, see subtype. If you specify a notificationSubtype, you need to also specify its related notificationType.
	NotificationSubtype notificationSubtype `json:"notificationSubtype"`

	// Optional. A Boolean value you set to true to request only the notifications that haven’t reached your server successfully.
	// The response also includes notifications that the App Store server is currently retrying to send to your server.
	OnlyFailures onlyFailures `json:"onlyFailures"`

	// Optional. The transaction identifier, which may be an original transaction identifier, of any transaction belonging to the customer.
	// Provide this field to limit the notification history request to this one customer.
	// Note: You may include either the transactionId or the notificationType property (or neither) in your query, but not both.
	TransactionId transactionId `json:"transactionId"`
}

// NotificationHistoryResponse A response that contains the App Store Server Notifications history for your app.
type NotificationHistoryResponse struct {
	// An array of App Store Server Notifications history records.
	// If you set onlyFailures to true in the NotificationHistoryRequest, this array contains only the notifications that failed to reach your server.
	NotificationHistory []notificationHistoryResponseItem `json:"notificationHistory"`

	// A Boolean value that indicates whether the App Store has more notification history records to send.
	// If hasMore is true, use the paginationToken in the subsequent request to get more records.
	// If hasMore is false, there are no more records available.
	HasMore hasMore `json:"hasMore"`

	// A pagination token that you provide to Get Notification History on a subsequent request to get the next page of responses.
	PaginationToken paginationToken `json:"paginationToken"`
}

// A unique identifier for a notification test that the App Store server sends to your server.
type testNotificationToken string

// SendTestNotificationResponse A response that contains the test notification token.
type SendTestNotificationResponse struct {
	// The test notification token that uniquely identifies the notification test that App Store Server Notifications sends to your server.
	TestNotificationToken testNotificationToken `json:"testNotificationToken"`
}

// CheckTestNotificationResponse A response that contains the contents of the App Store server’s test notification and the result from your server.
type CheckTestNotificationResponse struct {
	// An array of information the App Store server records for its attempts to send the TEST notification to your server. The array may contain a maximum of six sendAttemptItem objects.
	SendAttempts []sendAttemptItem `json:"sendAttempts"`

	// The signed payload, in JWS format, that contains the TEST notification that the App Store server sent to your server.
	SignedPayload signedPayload `json:"signedPayload"`

	// The result of the App Store server’s first attempt to send the TEST notification to your server.
	// Use the first sendAttemptItem in the sendAttempts array instead.
	FirstSendAttemptResult firstSendAttemptResult `json:"firstSendAttemptResult"`
}
