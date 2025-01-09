package models

// A notification type value that App Store Server Notifications 2 uses.
type notificationType string

// A notification subtype value that App Store Server Notifications 2 uses.
type notificationSubtype string

// A Boolean value that indicates whether the response includes only notifications that failed to reach your server.
type onlyFailures bool

// The unique identifier for a transaction, such as an In-App Purchase, restored In-App Purchase, or subscription renewal.
type transactionId string

// A pagination token that you return to the endpoint on a subsequent call to receive the next set of results.
type paginationToken string

// The success or error information the App Store server records when it attempts to send an App Store server notification to your server.
// SUCCESS: The App Store server received a success response when it sent the notification to your server.
// CIRCULAR_REDIRECT: The App Store server detected a continual redirect. Check your server’s redirects for a circular redirect loop.
// INVALID_RESPONSE: The App Store server received an invalid response from your server.
// NO_RESPONSE: The App Store server didn’t receive a valid HTTP response from your server.
// OTHER: Another error occurred that prevented your server from receiving the notification.
// PREMATURE_CLOSE: The App Store server’s connection to your server was closed while the send was in progress.
// SOCKET_ISSUE: A network error caused the notification attempt to fail.
// TIMED_OUT: The App Store server didn’t get a response from your server and timed out. Check that your server isn’t processing messages in line.
// TLS_ISSUE: The App Store server couldn’t establish a TLS session or validate your certificate. Check that your server has a valid certificate and supports Transport Layer Security (TLS) protocol 1.2 or later.
// UNSUCCESSFUL_HTTP_RESPONSE_CODE: The App Store server didn’t receive an HTTP 200 response from your server.
// UNSUPPORTED_CHARSET: The App Store server doesn’t support the supplied charset.
type sendAttemptResult string

type sendAttemptItem struct {
	// The date the App Store server attempts to send a notification.
	AttemptDate timestamp `json:"attemptDate"`
	// The success or error information the App Store server records when it attempts to send an App Store server notification to your server.
	SendAttemptResult sendAttemptResult `json:"sendAttemptResult"`
}
type signedPayload string
type firstSendAttemptResult string

// The App Store server notification history record,
// including the signed notification payload and the result of the server’s first send attempt.
type notificationHistoryResponseItem struct {
	// An array of information the App Store server records for its attempts to send a notification to your server.
	// The maximum number of entries in the array is six.
	SendAttempts []sendAttemptItem `json:"sendAttempts"`
	// The cryptographically signed payload, in JSON Web Signature (JWS) format,
	// containing the original response body of a version 2 notification. For more information,
	// see signedPayload in App Store Server Notifications.
	SignedPayload signedPayload `json:"signedPayload"`
	// The result of the App Store server’s first attempt to send the notification to your server’s App Store Server Notifications V2 endpoint.
	// Use the earliest sendAttemptItem in the sendAttempts array instead.
	FirstSendAttemptResult firstSendAttemptResult `json:"firstSendAttemptResult"`
}
