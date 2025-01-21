package ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

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
