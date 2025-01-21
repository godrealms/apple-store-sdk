package RequestTypes

import (
	"github.com/godrealms/apple-store-sdk/pkg/types"
)

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
	// A list of storefront country codes you provide to limit the storefronts that are eligible to receive the subscription-renewal-date extension.
	// Omit this list to request the subscription-renewal-date extension in all storefronts.
	StorefrontCountryCodes types.StorefrontCountryCodes `json:"storefrontCountryCodes"`
}
