package models

// The number of days to extend the subscription renewal date.
type extendByDays int32

// The code that represents the reason for the subscription-renewal-date extension.
// 0: Undeclared; no information provided.
// 1: The renewal-date extension is for customer satisfaction.
// 2: The renewal-date extension is for other reasons.
// 3: The renewal-date extension is due to a service issue or outage.
type extendReasonCode int32

// A string that contains a unique identifier you provide to track each subscription-renewal-date extension request.
type requestIdentifier string

type effectiveDate string

type webOrderLineItemId string

// The most recent App Store-signed transaction information and App Store-signed renewal information for an auto-renewable subscription.
type lastTransactions struct {
	// The original transaction identifier of the auto-renewable subscription.
	OriginalTransactionId originalTransactionId `json:"originalTransactionId"`

	// The status of the auto-renewable subscription.
	Status status `json:"status"`

	// The transaction information signed by the App Store, in JWS format.
	SignedTransactionInfo JWSTransaction `json:"signedTransactionInfo"`

	// The subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedRenewalInfo JWSRenewalInfo `json:"signedRenewalInfo"`
}

// SubscriptionGroupIdentifierItem Information for auto-renewable subscriptions,
// including signed transaction information and signed renewal information, for one subscription group.
type SubscriptionGroupIdentifierItem struct {
	// The subscription group identifier of the auto-renewable subscriptions in the lastTransactions array.
	SubscriptionGroupIdentifier subscriptionGroupIdentifier `json:"subscriptionGroupIdentifier"`

	// An array of the most recent App Store-signed transaction information and App Store-signed renewal information for all auto-renewable subscriptions in the subscription group.
	LastTransactions []lastTransactions `json:"lastTransactions"`
}

// The three-letter code that represents the country or region associated with the App Store storefront.
type storefrontCountryCode string

// A list of storefront country codes you provide to limit the storefronts for a subscription-renewal-date extension.
type storefrontCountryCodes []storefrontCountryCode

// A Boolean value that indicates whether the App Store completed the request to extend a subscription renewal date to active subscribers.
type complete bool

// The UNIX time, in milliseconds, that the App Store completes a request to extend a subscription renewal date for eligible subscribers.
type completeDate timestamp

// The count of subscriptions that fail to receive a subscription-renewal-date extension.
type failedCount int64

// The count of subscriptions that successfully receive a subscription-renewal-date extension.
type succeededCount int64
