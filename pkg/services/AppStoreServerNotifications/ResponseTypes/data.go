package responseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

type Data struct {
	// The unique identifier of the app that the notification applies to.
	// This property is available for apps that users download from the App Store.
	// It isn’t present in the sandbox environment.
	AppAppleId types.AppAppleId `json:"appAppleId"`

	// The bundle identifier of the app.
	BundleId types.BundleId `json:"bundleId"`

	// The version of the build that identifies an iteration of the bundle.
	BundleVersion types.BundleVersion `json:"bundleVersion"`

	// The reason the customer requested the refund.
	// This field appears only for CONSUMPTION_REQUEST notifications,
	// which the server sends when a customer initiates a refund request
	// for a consumable in-app purchase or auto-renewable subscription.
	ConsumptionRequestReason types.ConsumptionRequestReason `json:"consumptionRequestReason"`

	// The server environment that the notification applies to, either sandbox or production.
	Environment types.Environment `json:"environment"`

	// Subscription renewal information signed by the App Store,
	// in JSON Web Signature (JWS) format. This field appears only for notifications
	// that apply to auto-renewable subscriptions.
	SignedRenewalInfo types.JWSRenewalInfo `json:"signedRenewalInfo"`

	// Transaction information signed by the App Store, in JSON Web Signature (JWS) format.
	SignedTransactionInfo types.JWSTransaction `json:"signedTransactionInfo"`

	// The status of an auto-renewable subscription as of the signedDate in the responseBodyV2DecodedPayload.
	// This field appears only for notifications sent for auto-renewable subscriptions.
	Status types.Status `json:"status"`
}
