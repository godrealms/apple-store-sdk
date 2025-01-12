package models

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

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

// DecodedPayload Decrypt structure contents
func (sp signedPayload) DecodedPayload() (*NotificationsResponseBodyV2DecodedPayload, error) {
	// Delimiter information
	header, payloadBytes, signature, err := sp.parseSignedPayload()
	if err != nil {
		return nil, err
	}

	// Get public key information
	certificate, err := header.X5c.GetPublicKey()
	if err != nil {
		return nil, err
	}

	var payload NotificationsResponseBodyV2DecodedPayload
	if err = json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse payload JSON: %v", err)
	}

	singPayload := string(sp)
	signedContent := singPayload[:strings.LastIndex(singPayload, ".")]

	// Create a hash of the signed content
	hash := sha256.Sum256([]byte(signedContent))

	// Verify the signature
	err = rsa.VerifyPKCS1v15(certificate.PublicKey.(*rsa.PublicKey), crypto.SHA256, hash[:], signature)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %v", err)
	}

	return &payload, nil
}

// Parse signedPayload and return Header, Payload and Signature
func (sp signedPayload) parseSignedPayload() (*JWSDecodedHeader, []byte, []byte, error) {
	// Split signedPayload
	parts := strings.Split(string(sp), ".")
	if len(parts) != 3 {
		return nil, nil, nil, fmt.Errorf("invalid signedPayload format: expected 3 parts, got %d", len(parts))
	}

	// Decode Header
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode header: %w", err)
	}

	var header JWSDecodedHeader
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to unmarshal header: %w", err)
	}

	// 解码 Payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	// 解码 Signature
	signatureBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode signature: %w", err)
	}

	return &header, payloadBytes, signatureBytes, nil
}

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

// A string that provides details about select notification types in version 2.
// ACCEPTED
// Applies to the PRICE_INCREASE notificationType. A notification with this subtype indicates that the customer
// consented to the subscription price increase if the price increase requires customer consent, or that the system
// notified them of a price increase if the price increase doesn’t require customer consent.
//
// AUTO_RENEW_DISABLED
// Applies to the DID_CHANGE_RENEWAL_STATUS notificationType. A notification with this `subtype indicates that the user
// disabled subscription auto-renewal, or the App Store disabled subscription auto-renewal after the user requested a refund.
//
// AUTO_RENEW_ENABLED
// Applies to the DID_CHANGE_RENEWAL_STATUS notificationType. A notification with this subtype indicates that the user
// enabled subscription auto-renewal.
//
// BILLING_RECOVERY
// Applies to the DID_RENEW notificationType. A notification with this subtype indicates that the expired subscription
// that previously failed to renew has successfully renewed.
//
// BILLING_RETRY
// Applies to the EXPIRED notificationType. A notification with this subtype indicates that the subscription expired
// because the subscription failed to renew before the billing retry period ended.
//
// DOWNGRADE
// Applies to the DID_CHANGE_RENEWAL_PREF and OFFER_REDEEMED notificationType. A notification with this subtype indicates
// that the user downgraded their subscription or cross-graded to a subscription with a different duration.
// Downgrades take effect at the next renewal date.
//
// FAILURE
// Applies to the RENEWAL_EXTENSION notificationType. A notification with this subtype indicates that the
// subscription-renewal-date extension failed for an individual subscription. For details, see the data object in the
// responseBodyV2DecodedPayload. For information on the request, see Extend Subscription Renewal Dates for All Active Subscribers.
//
// GRACE_PERIOD
// Applies to the DID_FAIL_TO_RENEW notificationType. A notification with this subtype indicates that the subscription
// failed to renew due to a billing issue. Continue to provide access to the subscription during the grace period.
//
// INITIAL_BUY
// Applies to the SUBSCRIBED notificationType. A notification with this subtype indicates that the user purchased the
// subscription for the first time or that the user received access to the subscription through Family Sharing for the first time.
//
// PENDING
// Applies to the PRICE_INCREASE notificationType. A notification with this subtype indicates that the system informed
// the user of the subscription price increase, but the user hasn’t accepted it.
//
// PRICE_INCREASE
// Applies to the EXPIRED notificationType. A notification with this subtype indicates that the subscription expired
// because the user didn’t consent to a price increase.
//
// PRODUCT_NOT_FOR_SALE
// Applies to the EXPIRED notificationType. A notification with this subtype indicates that the subscription expired
// because the product wasn’t available for purchase at the time the subscription attempted to renew.
//
// RESUBSCRIBE
// Applies to the SUBSCRIBED notificationType. A notification with this subtype indicates that the user resubscribed or
// received access through Family Sharing to the same subscription or to another subscription within the same subscription group.
//
// SUMMARY
// Applies to the RENEWAL_EXTENSION notificationType. A notification with this subtype indicates that the App Store server
// completed your request to extend the subscription renewal date for all eligible subscribers. For the summary details,
// see the summary object in the responseBodyV2DecodedPayload. For information on the request,
// see Extend Subscription Renewal Dates for All Active Subscribers.
//
// UPGRADE
// Applies to the DID_CHANGE_RENEWAL_PREF and OFFER_REDEEMED notificationType.
// A notification with this subtype indicates that the user upgraded their subscription or cross-graded to a subscription
// with the same duration. Upgrades take effect immediately.
//
// UNREPORTED
// Applies to the EXTERNAL_PURCHASE_TOKEN notificationType.
// A notification with this subtype indicates that Apple created a token for your app but didn’t receive a report.
// For more information about reporting the token, see externalPurchaseToken.
//
// VOLUNTARY
// Applies to the EXPIRED notificationType.
// A notification with this subtype indicates that the subscription expired after the user disabled subscription auto-renewal.
type subtype string

// The version of the build that identifies an iteration of the bundle.
type bundleVersion string

// The customer-provided reason for a refund request.
// UNINTENDED_PURCHASE: The customer didn’t intend to make the in-app purchase.
// FULFILLMENT_ISSUE: The customer had issues with receiving or using the in-app purchase.
// UNSATISFIED_WITH_PURCHASE: The customer wasn’t satisfied with the in-app purchase.
// LEGAL: The customer requested a refund based on a legal reason.
// OTHER: The customer requested a refund for other reasons.
type consumptionRequestReason string

// The payload data that contains app metadata and the signed renewal and transaction information.
type data struct {
	// The unique identifier of the app that the notification applies to.
	// This property is available for apps that users download from the App Store.
	// It isn’t present in the sandbox environment.
	AppAppleId appAppleId `json:"appAppleId"`

	// The bundle identifier of the app.
	BundleId bundleId `json:"bundleId"`

	// The version of the build that identifies an iteration of the bundle.
	BundleVersion bundleVersion `json:"bundleVersion"`

	// The reason the customer requested the refund. This field appears only for CONSUMPTION_REQUEST notifications,
	//which the server sends when a customer initiates a refund request for a consumable in-app purchase or auto-renewable subscription.
	ConsumptionRequestReason consumptionRequestReason `json:"consumptionRequestReason"`

	// The server environment that the notification applies to, either sandbox or production.
	Environment environment `json:"environment"`

	// Subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format.
	// This field appears only for notifications that apply to auto-renewable subscriptions.
	SignedRenewalInfo JWSRenewalInfo `json:"signedRenewalInfo"`

	// Transaction information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedTransactionInfo JWSTransaction `json:"signedTransactionInfo"`

	// The status of an auto-renewable subscription as of the signedDate in the responseBodyV2DecodedPayload.
	// This field appears only for notifications sent for auto-renewable subscriptions.
	Status status `json:"status"`
}

type summary struct {
	// The UUID that represents a specific request to extend a subscription renewal date.
	// This value matches the value you initially specify in the requestIdentifier when you call Extend Subscription
	// Renewal Dates for All Active Subscribers in the App Store Server API.
	RequestIdentifier requestIdentifier `json:"requestIdentifier"`

	// The server environment that the notification applies to, either sandbox or production.
	Environment environment `json:"environment"`

	// The unique identifier of the app that the notification applies to. This property is available for apps that users
	// download from the App Store. It isn’t present in the sandbox environment.
	AppAppleId appAppleId `json:"appAppleId"`

	// The bundle identifier of the app.
	BundleId bundleId `json:"bundleId"`

	// The product identifier of the auto-renewable subscription that the subscription-renewal-date extension applies to.
	ProductId productId `json:"productId"`

	// A list of country codes that limits the App Store’s attempt to apply the subscription-renewal-date extension.
	// If this list isn’t present, the subscription-renewal-date extension applies to all storefronts.
	StorefrontCountryCodes storefrontCountryCodes `json:"storefrontCountryCodes"`

	// The final count of subscriptions that fail to receive a subscription-renewal-date extension.
	FailedCount failedCount `json:"failedCount"`

	// The final count of subscriptions that successfully receive a subscription-renewal-date extension.
	SucceededCount succeededCount `json:"succeededCount"`
}

// The field of an external purchase token that uniquely identifies the token.
type externalPurchaseId string

// The field of an external purchase token that uniquely identifies the token.
type tokenCreationDate timestamp

// The payload data that contains an external purchase token.
type externalPurchaseToken struct {
	// The unique identifier of the token. Use this value to report tokens and their associated transactions in the Send External Purchase Report endpoint.
	ExternalPurchaseId externalPurchaseId `json:"externalPurchaseId"`

	// The UNIX time, in milliseconds, when the system created the token.
	TokenCreationDate tokenCreationDate `json:"tokenCreationDate"`

	// The app Apple ID for which the system generated the token.
	AppAppleId appAppleId `json:"appAppleId"`

	// The bundle ID of the app for which the system generated the token.
	BundleId bundleId `json:"bundleId"`
}

// A string that indicates the notification’s App Store Server Notifications version number.
type version string

// A unique identifier for the notification.
type notificationUUID string
