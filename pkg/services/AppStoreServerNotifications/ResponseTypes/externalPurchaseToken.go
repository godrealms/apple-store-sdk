package responseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

// The payload data that contains an external purchase token.
type externalPurchaseToken struct {
	// The unique identifier of the token. Use this value to report tokens and their associated transactions in the Send External Purchase Report endpoint.
	ExternalPurchaseId types.ExternalPurchaseId `json:"externalPurchaseId"`

	// The UNIX time, in milliseconds, when the system created the token.
	TokenCreationDate types.Timestamp `json:"tokenCreationDate"`

	// The app Apple ID for which the system generated the token.
	AppAppleId types.AppAppleId `json:"appAppleId"`

	// The bundle ID of the app for which the system generated the token.
	BundleId types.BundleId `json:"bundleId"`
}
