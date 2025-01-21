package models

import "github.com/godrealms/apple-store-sdk/pkg/types"

// StatusResponse A response that contains status information for all of a customer’s auto-renewable subscriptions in your app.
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

// ExtendRenewalDateRequest The request body that contains subscription-renewal-extension data for an individual subscription.
type ExtendRenewalDateRequest struct {
	//Required. The number of days to extend the subscription renewal date.
	//Maximum: 90
	ExtendByDays types.ExtendByDays `json:"extendByDays"`

	//Required. The reason code for the subscription date extension.
	ExtendReasonCode types.ExtendReasonCode `json:"extendReasonCode"`

	//Required. A string that contains a value you provide to uniquely identify this renewal-date extension request.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`
}

// ExtendRenewalDateResponse A response that indicates whether an individual renewal-date extension succeeded, and related details.
type ExtendRenewalDateResponse struct {
	// The new subscription expiration date of a successful subscription-renewal-date extension.
	EffectiveDate types.EffectiveDate `json:"effectiveDate"`

	// The original transaction identifier of the subscription that experienced a service interruption.
	OriginalTransactionId types.OriginalTransactionId `json:"originalTransactionId"`

	// A Boolean value that indicates whether the subscription-renewal-date extension succeeded.
	Success types.Success `json:"success"`

	// A unique ID that identifies subscription-purchase events, including subscription renewals, across devices.
	WebOrderLineItemId types.WebOrderLineItemId `json:"webOrderLineItemId"`
}

// MassExtendRenewalDateRequest The request body that contains subscription-renewal-extension data to apply for all eligible active subscribers.
type MassExtendRenewalDateRequest struct {
	// Required. A string that contains a one-time UUID value you provide to identify this subscription-renewal-date extension request.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`

	//Required. The number of days to extend the subscription renewal date.
	//Maximum: 90
	ExtendByDays types.ExtendByDays `json:"extendByDays"`

	// Required. The reason code for the subscription-renewal-date extension.
	ExtendReasonCode types.ExtendReasonCode `json:"extendReasonCode"`

	// Required. The product identifier of the auto-renewable subscription that you’re requesting the renewal-date extension for.
	ProductId types.ProductId `json:"productId"`

	// A list of storefront country codes you provide to limit the storefronts that are eligible to receive the subscription-renewal-date extension. Omit this list to request the subscription-renewal-date extension in all storefronts.
	StorefrontCountryCodes types.StorefrontCountryCodes `json:"storefrontCountryCodes"`
}

// MassExtendRenewalDateResponse A response that indicates the server successfully received the subscription-renewal-date extension request.
type MassExtendRenewalDateResponse struct {
	// A string that contains the UUID that identifies the subscription-renewal-date extension request.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`
}

// MassExtendRenewalDateStatusResponse A response that indicates the current status of a request to extend the subscription renewal date to all eligible subscribers.
type MassExtendRenewalDateStatusResponse struct {
	//The UUID that represents your request for a subscription-renewal-date extension.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`

	//A Boolean value that’s TRUE to indicate that the App Store completed your request to extend a subscription renewal date for all eligible subscribers.
	//The value is FALSE if the request is in progress.
	Complete types.Complete `json:"complete"`

	//The date that the App Store completes the request.
	//Appears only if complete is TRUE.
	CompleteDate types.Timestamp `json:"completeDate"`

	//The final count of subscribers that fail to receive a subscription-renewal-date extension.
	//Appears only if complete is TRUE.
	FailedCount types.FailedCount `json:"failedCount"`

	//The final count of subscribers that successfully receive a subscription-renewal-date extension.
	//Appears only if complete is TRUE.
	SucceededCount types.SucceededCount `json:"succeededCount"`
}
