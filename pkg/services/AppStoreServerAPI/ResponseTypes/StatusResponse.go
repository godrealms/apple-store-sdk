package ResponseTypes

import (
	"github.com/godrealms/apple-store-sdk/pkg/types"
)

// LastTransactionsItem
// The most recent App Store-signed transaction information and App Store-signed renewal information for an auto-renewable subscription.
type LastTransactionsItem struct {
	// The original transaction identifier of the auto-renewable subscription.
	OriginalTransactionId types.OriginalTransactionId `json:"originalTransactionId"`
	// The status of the auto-renewable subscription.
	Status types.Status `json:"status"`
	// The subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedRenewalInfo types.JWSRenewalInfo `json:"signedRenewalInfo"`
	// The transaction information signed by the App Store, in JWS format.
	SignedTransactionInfo types.JWSTransaction `json:"signedTransactionInfo"`
}

// SubscriptionGroupIdentifierItem
// Information for auto-renewable subscriptions,
// including signed transaction information and signed renewal information,
// for one subscription group.
type SubscriptionGroupIdentifierItem struct {
	// The subscription group identifier of the auto-renewable subscriptions in the lastTransactions array.
	SubscriptionGroupIdentifier types.SubscriptionGroupIdentifier `json:"subscriptionGroupIdentifier"`
	// An array of the most recent App Store-signed transaction information and App Store-signed renewal
	// information for all auto-renewable subscriptions in the subscription group.
	LastTransactions []LastTransactionsItem `json:"lastTransactions"`
}

// StatusResponse
// A response that contains status information for all of a customer’s auto-renewable subscriptions in your app.
type StatusResponse struct {
	// An array of information for auto-renewable subscriptions, including App Store-signed transaction information and App Store-signed renewal information.
	Data []SubscriptionGroupIdentifierItem `json:"data"`
	// The server environment, sandbox or production, in which the App Store generated the response.
	Environment types.Environment `json:"environment"`
	// Your app’s App Store identifier.
	AppAppleId types.AppAppleId `json:"appAppleId"`
	// Your app’s bundle identifier.
	BundleId types.BundleId `json:"bundleId"`
}
