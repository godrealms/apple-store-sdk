package _ResponseTypes

// NotificationsResponseBodyV2 The response body the App Store sends in a version 2 server notification.
type NotificationsResponseBodyV2 struct {
	// A cryptographically signed payload, in JSON Web Signature (JWS) format, that contains the response body for a version 2 notification.
	SignedPayload signedPayload `json:"signedPayload"`
}
