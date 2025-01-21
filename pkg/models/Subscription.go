package models

import "github.com/godrealms/apple-store-sdk/pkg/types"

// The most recent App Store-signed transaction information and App Store-signed renewal information for an auto-renewable subscription.
type lastTransactions struct {
	// The original transaction identifier of the auto-renewable subscription.
	OriginalTransactionId types.OriginalTransactionId `json:"originalTransactionId"`

	// The status of the auto-renewable subscription.
	Status types.Status `json:"status"`

	// The transaction information signed by the App Store, in JWS format.
	SignedTransactionInfo types.JWSTransaction `json:"signedTransactionInfo"`

	// The subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedRenewalInfo types.JWSRenewalInfo `json:"signedRenewalInfo"`
}

// SubscriptionGroupIdentifierItem Information for auto-renewable subscriptions,
// including signed transaction information and signed renewal information, for one subscription group.
type SubscriptionGroupIdentifierItem struct {
	// The subscription group identifier of the auto-renewable subscriptions in the lastTransactions array.
	SubscriptionGroupIdentifier types.SubscriptionGroupIdentifier `json:"subscriptionGroupIdentifier"`

	// An array of the most recent App Store-signed transaction information and App Store-signed renewal information for all auto-renewable subscriptions in the subscription group.
	LastTransactions []lastTransactions `json:"lastTransactions"`
}
